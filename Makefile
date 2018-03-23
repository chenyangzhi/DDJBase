all: build

fmt:
	gofmt -l -w -s src/

test:


build:dep
	go build -o bin/server src/main.go
	go build -o bin/grpc_client src/grpc_client.go
	go build -o bin/grpc_server src/grpc_server.go

clean:
	rm -rf output

dep:fmt
