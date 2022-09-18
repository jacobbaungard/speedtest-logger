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
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config struct {
	LibrespeedBinary string
	CronSpec         string
	InfluxHost       string
	InfluxPort       string
	InfluxToken      string // v2
	InfluxOrg        string // v2
	InfluxBucket     string // v2
	InfluxUsername   string // v1
	InfluxPassword   string // v1
	InfluxDatabase   string // v1
	InfluxSSL        bool
}

func ParseConfig(ConfigFile string) Config {
	// Set defaults
	viper.SetDefault("LibrespeedBinary", "/usr/bin/librespeed-cli")
	viper.SetDefault("CronSpec", "0 * * * *")
	viper.SetDefault("InfluxHost", "localhost")
	viper.SetDefault("InfluxPort", "8086")
	viper.SetDefault("InfluxSSL", true)

	// Read config
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(ConfigFile), filepath.Ext(ConfigFile)))
	viper.AddConfigPath(filepath.Dir(ConfigFile))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Build config struct
	var conf Config
	conf.LibrespeedBinary = viper.GetString("LibrespeedBinary")
	conf.CronSpec = viper.GetString("CronSpec")
	conf.InfluxHost = viper.GetString("InfluxHost")
	conf.InfluxPort = viper.GetString("InfluxPort")
	conf.InfluxToken = viper.GetString("InfluxToken")
	conf.InfluxOrg = viper.GetString("InfluxOrg")
	conf.InfluxBucket = viper.GetString("InfluxBucket")
	conf.InfluxUsername = viper.GetString("InfluxUsername")
	conf.InfluxPassword = viper.GetString("InfluxPassword")
	conf.InfluxDatabase = viper.GetString("InfluxDatabase")
	conf.InfluxSSL = viper.GetBool("InfluxSSL")

	// validate binary
	// validate influx v1/v1
	// validate cron spec

	return conf

}
