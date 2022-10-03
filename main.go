/*
Copyright © 2022 Jacob Baungard Hansen me@jacobbaungard.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"flag"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jacobbaungard/speedtest-logger/internal"
	"github.com/rs/zerolog/log"
)

func main() {

	// init logger
	internal.InitializeLogger()
	// Load config
	var configfile = flag.String("config", "/etc/speedtest-logger.yaml", "Location of the configuration file")
	flag.Parse()

	// Parse config
	conf := internal.ParseConfig(*configfile)
	internal.SetLogLevel(conf.LogLevel)
	if !internal.ValidateConfig(conf) {
		log.Fatal().Msg("Config validation failed")
	}

	s := gocron.NewScheduler(time.UTC)

	_, err := s.Cron(conf.CronSpec).Do(func() { internal.Run(conf) })
	if err != nil {
		log.Fatal().Err(err).Msg("Error scheduling task")
	}

	log.Info().Msg("speedtest-logger started. Awaiting next run")
	s.StartBlocking()
}
