include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	set GOBIN=$(LOCAL_BIN) && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	set GOBIN=$(LOCAL_BIN) && go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	protoc --proto_path api/sum_v1 \
	--go_out=pkg/sum_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/sum_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/sum_v1/user.proto