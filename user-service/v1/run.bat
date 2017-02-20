@echo off
set USER_SERVICE_NAME=UserService
set USER_SERVICE_VERSION=v1
set USER_SERVICE_PORT=7777
set DISCOVERY=http://localhost:9999
set DATABASE_DIALECT=postgres
set DATABASE=host=localhost user=postgres dbname=user_service sslmode=disable password=
set UPGRADE_DATABASE=1
rem export METRICS="http://localhost:8086"
rem export METRICS_DATABASE="metrics"
rem export TRACER="http://localhost:9411/api/v1/spans"

go run main.go