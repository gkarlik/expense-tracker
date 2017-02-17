export USER_SERVICE_NAME="UserService"
export USER_SERVICE_VERSION="v1"
export USER_SERVICE_PORT="7777"
export DISCOVERY="http://localhost:9999"
export DATABASE_DIALECT="postgres"
export DATABASE="host=localhost user=postgres dbname=user_service sslmode=disable password="
export UPGRADE_DATABASE="1"
# export METRICS="http://localhost:8086"
# export METRICS_DATABASE="metrics"
# export TRACER="http://localhost:9411/api/v1/spans"

go run main.go