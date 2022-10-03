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

package internal

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	LibrespeedBinary string
	CronSpec         string
	LogLevel         string
	InfluxAddress    string
	InfluxToken      string // v2
	InfluxOrg        string // v2
	InfluxBucket     string // v2
	InfluxUsername   string // v1
	InfluxPassword   string // v1
	InfluxDatabase   string // v1
	InfluxSSL        bool   // ?
}

func ParseConfig(ConfigFile string) Config {
	// Set defaults
	viper.SetDefault("LibrespeedBinary", "/usr/bin/librespeed-cli")
	viper.SetDefault("CronSpec", "0 * * * *")
	viper.SetDefault("InfluxAddress", "http://localhost:8086")
	viper.SetDefault("InfluxSSL", true)
	viper.SetDefault("LogLevel", "info")

	// Read config
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(ConfigFile), filepath.Ext(ConfigFile)))
	viper.AddConfigPath(filepath.Dir(ConfigFile))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal().Err(err).Msg("Error reading config file")
	}

	// Build config struct
	var conf Config
	conf.LibrespeedBinary = viper.GetString("LibrespeedBinary")
	conf.CronSpec = viper.GetString("CronSpec")
	conf.InfluxAddress = viper.GetString("InfluxAddress")
	conf.InfluxToken = viper.GetString("InfluxToken")
	conf.InfluxOrg = viper.GetString("InfluxOrg")
	conf.InfluxBucket = viper.GetString("InfluxBucket")
	conf.InfluxUsername = viper.GetString("InfluxUsername")
	conf.InfluxPassword = viper.GetString("InfluxPassword")
	conf.InfluxDatabase = viper.GetString("InfluxDatabase")
	conf.InfluxSSL = viper.GetBool("InfluxSSL")
	conf.LogLevel = viper.GetString("LogLevel")

	return conf

}

func ValidateConfig(conf Config) bool {
	// Check if librespeed binary exists
	if _, err := os.Stat(conf.LibrespeedBinary); errors.Is(err, os.ErrNotExist) {
		log.Err(err).Msg("Could not find librespeed binary")
		return false
	}

	// Check for valid cron spec
	if _, err := cron.ParseStandard(conf.CronSpec); err != nil {
		log.Err(err).Msg("Unable to parse cron spec")
		return false
	}

	// Check address is sensible
	if _, err := url.ParseRequestURI(conf.InfluxAddress); err != nil {
		log.Err(err).Msg("Invalid Influx address")
		return false
	}

	// Cannot set both InfluxToken and InfluxUsername/Password
	if (len(conf.InfluxUsername) > 0 || len(conf.InfluxPassword) > 0) && len(conf.InfluxToken) > 0 {
		log.Error().Msg("Cannot use both InfluxUsername and InfluxToken. Use username for Influx v1, token for Influx v2")
		return false
	}

	// If using InfluxV1 we need to ensure both username and password is set
	if len(conf.InfluxUsername) > 0 || len(conf.InfluxPassword) > 0 {
		if len(conf.InfluxUsername) == 0 || len(conf.InfluxPassword) == 0 {
			log.Error().Msg("If using Influx V1, both a username and a password is needed")
			return false
		}
	}

	return true
}
