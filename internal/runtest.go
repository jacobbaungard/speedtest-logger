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
	"encoding/json"
	"fmt"
	"github.com/commander-cli/cmd"
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
	fmt.Println("Starting speedtest")
	c := cmd.NewCommand("librespeed-cli --json")
	err := c.Execute()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(c.Stdout())
	// Parse result/jsoni
	var results []Result
	err = json.Unmarshal([]byte(c.Stdout()), &results)

	if err != nil {
		fmt.Println("Unable to parse JSON.")
		fmt.Println(c.Stdout())
		panic(err.Error())
	}

	fmt.Println("Struct is:", results[0])
	// save to influx
	WriteResult(conf, results[0])
}
