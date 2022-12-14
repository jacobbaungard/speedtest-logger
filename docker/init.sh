#!/bin/bash
CONFIG_FILE="/etc/speedtest-logger.yaml"

cat <<EOF > $CONFIG_FILE
---
LibrespeedBinary: $LIBRESPEED_BINARY
LibrespeedServer: $LIBRESPEED_SERVER
InfluxAddress: $INFLUX_ADDRESS
InfluxOrg: $INFLUX_ORG
InfluxBucket: $INFLUX_BUCKET
InfluxToken: $INFLUX_TOKEN
InfluxUsername: $INFLUX_USERNAME
InfluxPassword: $INFLUX_PASSWORD
InfluxDatabase: $INFLUX_DATABASE
CronSpec: $CRON_SPEC
LogLevel: $LOG_LEVEL
EOF

$@
