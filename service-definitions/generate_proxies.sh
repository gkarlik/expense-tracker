# regenerate proxies for protocol buffer definitions
protoc -I protos protos/user-service/v1/user-service.proto --go_out=plugins=grpc:proxies
cp -r proxies/user-service/v1/* ../user-service/v1/proxy/
cp -r proxies/user-service/v1/* ../api-gateway/proxy/user-service/v1

protoc -I protos protos/expense-service/v1/expense-service.proto --go_out=plugins=grpc:proxies
cp -r proxies/expense-service/v1/* ../expense-service/v1/proxy/