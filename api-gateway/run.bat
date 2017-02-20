@echo off
set GATEWAY_NAME=ApiGateway
set GATEWAY_VERSION=v1
set GATEWAY_PORT=9999
set GATEWAY_SECRET=this is secret
set DISCOVERY=localhost:9999
rem export METRICS="http://localhost:8086"
rem export METRICS_DATABASE="metrics"
rem export TRACER="http://localhost:9411/api/v1/spans"

go run main.go