#!/bin/bash
CONFIG_FILE="/etc/speedtest-logger.yaml"

echo "running init"

if [ -f "$CONFIG_FILE" ]; then
  echo "a config file was already found"
  exit 0
fi

cat <<EOF > $CONFIG_FILE
LibrespeedBinary: $LIBRESPEED_BINARY
InfluxHost: $INFLUX_HOST
InfluxPort: $INFLUX_PORT
InfluxOrg: $INFLUX_ORG
InfluxBucket: $INFLUX_BUCKET
InfluxToken: $INFLUX_TOKEN
CronSpec: $CRON_SPEC
EOF

echo "Wrote config file"

$@
