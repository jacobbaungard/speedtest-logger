# Speedtest logger
![GitHub Workflow Status (event)](https://img.shields.io/github/workflow/status/jacobbaungard/speedtest-logger/CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/jacobbaungard/speedtest-logger)](https://goreportcard.com/report/github.com/jacobbaungard/speedtest-logger)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/jacobbaungard/speedtest-logger)
![GitHub](https://img.shields.io/github/license/jacobbaungard/speedtest-logger)

Speedtest logger runs an internet speedtest using
[Librespeed](https://librespeed.org/) at a regular interval and saves it to
InfluxDB.

## Usage

### Docker-compose

```
version: '3'
services:
  speedtest-logger:
     image: jacobbaungard/speedtest-logger
     container_name: speedtest-logger
     environment:
        - INFLUX_ADDRESS="http://localhost:8086"
        - INFLUX_ORG="ORG"
        - INFLUX_BUCKET="BUCKET"
        - INFLUX_TOKEN="TOKEN"
        - CRON_SPEC="0 * * * *"
```

### Configuration

The following configuration options are available.

|     Setting               |   Description                                                  |
|     :----:                |   ---                                                          |
|     `CRON_SPEC`           |   When to run the speedtest in regular cron format             |
|     `INFLUX_ADDRESS`      |   Address to InfluxDB server, example `http://localhost:8086`  |
|     `INFLUX_ORG`          |   InfluxDB Organisation ID (InfluxDB 2)                        |
|     `INFLUX_BUCKET`       |   InfluxDB Bucket ID (InfluxDB 2)                              |
|     `INFLUX_TOKEN`        |   InfluxDB Authentication Token (InfluxDB 2)                   |
|     `LOG_LEVEL`           |   Log level: debug, info (default), warn, error                |
|     `LIBRESPEED_SERVER`   |   By default a server is chosen based on lowest ping. Use this option to use a specific for the speedtests. A list of server IDs can be found [here](https://librespeed.org/backend-servers/servers.php).          |

## Similar tools
- [Speedtest-Tracker](https://github.com/henrywhitaker3/Speedtest-Tracker)
- [Speedtest to InfluxDB](https://github.com/aidengilmartin/speedtest-to-influxdb)
