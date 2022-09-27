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
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func WriteResult(conf Config, result Result) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	// TODO: v1?
	client := influxdb2.NewClient(conf.InfluxAddress, conf.InfluxToken)
	if client == nil {
		fmt.Println("Client was null")
		return
	}
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(conf.InfluxOrg, conf.InfluxBucket)
	if writeAPI == nil {
		fmt.Println("WriteAPI was null")
		return
	}
	p := influxdb2.NewPointWithMeasurement("speedtest").
		AddTag("server", result.Server.Name).
		AddField("bytes_sent", result.BytesSent).
		AddField("bytes_recieved", result.BytesRecieved).
		AddField("ping", result.Ping).
		AddField("jitter", result.Jitter).
		AddField("download", result.Download).
		AddField("upload", result.Upload).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client.Close()
	fmt.Println("Wrote to influx")
}
