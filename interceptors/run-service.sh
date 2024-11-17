#!/bin/bash
set -e

echo "Building interceptors service..."
go build -o bin/interceptors ./cmd/interceptors
go build -o bin/interceptors-cli ./cmd/interceptors-cli

echo "Starting interceptors service..."
./bin/interceptors --http-port=8088 &
SERVICE_PID=$!

# Ensure we kill the service when the script exits
trap "kill $SERVICE_PID" EXIT

echo "Service is running on http://localhost:8088"
echo "Press Ctrl+C to stop the service"

# Keep the script running
wait $SERVICE_PID
