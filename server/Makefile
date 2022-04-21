.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service_proto/service.proto

.PHONY: swagger-ui 
swagger-ui:
	docker run -p 8080:8080 -e SWAGGER_JSON=/tmp/service.swagger.json -v `pwd`/swagger/service_proto:/tmp swaggerapi/swagger-ui
