#!/usr/bin/env bash

host="$1"
port="$2"

echo "Remote forwarding: $host:/dev/ttyUSB0 -> $host:$port"
ssh "ubuntu@$host" "socat -s /dev/ttyUSB0,raw,echo=0 tcp-listen:$port,reuseaddr" &
pid="$!"

function cleanup {
  echo "Killing remote forwarding..."
  kill "$pid"
}
trap cleanup EXIT

echo "Waiting for remote forwarding..."
sleep 5

echo "Local forwarding: $host:$port -> localhost:dev/ttyUSB0"
socat -d -v PTY,raw,echo=0,link=dev/ttyUSB0 "tcp:$host:$port"
