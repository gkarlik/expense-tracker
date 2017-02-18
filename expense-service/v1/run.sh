export EXPENSE_SERVICE_NAME="ExpenseService"
export EXPENSE_SERVICE_VERSION="v1"
export EXPENSE_SERVICE_PORT="5555"
export DISCOVERY="http://localhost:9999"
export DATABASE_DIALECT="postgres"
export DATABASE="host=localhost user=postgres dbname=expense_service sslmode=disable password="
export UPGRADE_DATABASE="1"
# export METRICS="http://localhost:8086"
# export METRICS_DATABASE="metrics"
# export TRACER="http://localhost:9411/api/v1/spans"

go run main.go