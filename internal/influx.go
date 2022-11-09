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
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog/log"
	"time"
)

func WriteResult(conf Config, result Result) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	var client influxdb2.Client
	var InfluxOrg string
	var InfluxBucket string

	// assume v2 if token is provided, otherwise we assume Influx 1.8+
	if len(conf.InfluxToken) > 0 {
		client = influxdb2.NewClient(conf.InfluxAddress, conf.InfluxToken)
		InfluxBucket = conf.InfluxBucket
		InfluxOrg = conf.InfluxOrg
	} else {
		client = influxdb2.NewClient(conf.InfluxAddress, conf.InfluxUsername+":"+conf.InfluxPassword)
		InfluxBucket = conf.InfluxDatabase
		InfluxOrg = ""
	}
	if client == nil {
		log.Error().Msg("Error initializing influx client library")
		return
	}
	// Use blocking write client for writes to desired bucket
	log.Debug().Str("InfluxDB bucket/database", InfluxBucket).Msg("Setting up InfluxDB client")
	writeAPI := client.WriteAPIBlocking(InfluxOrg, InfluxBucket)
	if writeAPI == nil {
		log.Error().Msg("Error initializing influx WriteAPI")
		return
	}
	p := influxdb2.NewPointWithMeasurement("speedtest").
		AddTag("server", result.Server.Name).
		AddField("bytes_sent", result.BytesSent).
		AddField("bytes_received", result.BytesReceived).
		AddField("ping", result.Ping).
		AddField("jitter", result.Jitter).
		AddField("download", result.Download).
		AddField("upload", result.Upload).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Err(err).Msg("Error writing to influx")
		return
	}
	client.Close()
	log.Debug().Msg("Wrote datapoint to InfluxDB")
}
