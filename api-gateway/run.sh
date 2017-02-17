export GATEWAY_NAME="ApiGateway"
export GATEWAY_VERSION="v1"
export GATEWAY_PORT="9999"
export GATEWAY_SECRET="this is secret"
export DISCOVERY="localhost:9999"
# export METRICS="http://localhost:8086"
# export METRICS_DATABASE="metrics"
# export TRACER="http://localhost:9411/api/v1/spans"

go run main.go