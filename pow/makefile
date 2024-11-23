include .env
LOCAL_BIN:=$(CURDIR)/bin
MICROSERVICE:=pow

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir .\pkg\$(MICROSERVICE)_v1
	protoc --proto_path api/$(MICROSERVICE)_v1 --go_out=pkg/$(MICROSERVICE)_v1 --go_opt=paths=source_relative --plugin=protoc-gen-go.exe --go-grpc_out=pkg/$(MICROSERVICE)_v1 --go-grpc_opt=paths=source_relative --plugin=protoc-gen-go-grpc.exe api/$(MICROSERVICE)_v1/$(MICROSERVICE).proto