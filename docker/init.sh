#!/bin/bash
CONFIG_FILE="/etc/speedtest-logger.yaml"

echo "running init"

if [ -f "$CONFIG_FILE" ]; then
  echo "a config file was already found"
else

cat <<EOF > $CONFIG_FILE
---
LibrespeedBinary: $LIBRESPEED_BINARY
InfluxHost: $INFLUX_HOST
InfluxPort: $INFLUX_PORT
InfluxOrg: $INFLUX_ORG
InfluxBucket: $INFLUX_BUCKET
InfluxToken: $INFLUX_TOKEN
CronSpec: $CRON_SPEC
EOF
  
  echo "Wrote new config file"
fi

$@
