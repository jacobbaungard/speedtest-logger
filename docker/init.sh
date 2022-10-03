#!/bin/bash
CONFIG_FILE="/etc/speedtest-logger.yaml"

echo "running init"

if [ -f "$CONFIG_FILE" ]; then
  echo "a config file was already found"
else

cat <<EOF > $CONFIG_FILE
---
LibrespeedBinary: $LIBRESPEED_BINARY
InfluxAddress: $INFLUX_ADDRESS
InfluxOrg: $INFLUX_ORG
InfluxBucket: $INFLUX_BUCKET
InfluxToken: $INFLUX_TOKEN
CronSpec: $CRON_SPEC
LogLevel: $LOG_LEVEL
EOF

  echo "Wrote new config file"
fi

$@
