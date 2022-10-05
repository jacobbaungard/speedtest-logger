/*
Copyright © 2022 The speedtest-logger developers

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
	"encoding/json"
	"github.com/commander-cli/cmd"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Result struct {
	Timestamp     string  `json:"timestamp"` //TODO: check type?
	Server        Server  `json:"server"`
	Client        Client  `json:"client"`
	BytesSent     int     `json:"bytes_sent"`
	BytesRecieved int     `json:"bytes_received"`
	Ping          float64 `json:"ping"`
	Jitter        float64 `json:"jitter"`
	Upload        float64 `json:"upload"`
	Download      float64 `json:"download"`
	Share         string  `json:"share"`
}

type Server struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Client struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func Run(conf Config) {

	// Run run_test
	log.Info().Msg("Starting speedtest")

	librespeedCommand := "librespeed-cli --json"
	if conf.LibrespeedServer != -1 {
		librespeedCommand += " --server " + strconv.Itoa(conf.LibrespeedServer)
	}
	log.Debug().Str("Librespeed command", librespeedCommand).Msg("Constructed Librespeed command")

	c := cmd.NewCommand(librespeedCommand)
	err := c.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("Error executing librespeed-cli binary")
	}

	log.Debug().RawJSON("librespeed-cli raw output", []byte(c.Stdout())).Msg("librespeed-cli call completed")

	// Parse result/jsoni
	var results []Result
	err = json.Unmarshal([]byte(c.Stdout()), &results)

	if err != nil {
		log.Fatal().
			RawJSON("librespeed-cli raw output", []byte(c.Stdout())).Err(err).
			Msg("Unable to parse JSON from librespeed-cli")
	}

	// Convert struct back to json for debug printing (if enabled)
	if e := log.Debug(); e.Enabled() {
		resultJSON, err := json.Marshal(results[0])
		if err != nil {
			log.Warn().Err(err).Msg("Unable to output struct to JSON")
		} else {
			log.Debug().RawJSON("Parsed results", resultJSON).Msg("Speedtest result")
		}
	}

	// save to influx
	WriteResult(conf, results[0])
	log.Info().
		Float64("Download", results[0].Download).
		Float64("Upload", results[0].Upload).
		Msg("Speedtest complete.")
}
