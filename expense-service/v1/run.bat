@echo off
set EXPENSE_SERVICE_NAME=ExpenseService
set EXPENSE_SERVICE_VERSION=v1
set EXPENSE_SERVICE_PORT=5555
set DISCOVERY=http://localhost:9999
set DATABASE_DIALECT=postgres
set DATABASE=host=localhost user=postgres dbname=expense_service sslmode=disable password=
set UPGRADE_DATABASE=1
rem export METRICS="http://localhost:8086"
rem export METRICS_DATABASE="metrics"
rem export TRACER="http://localhost:9411/api/v1/spans"

go run main.go