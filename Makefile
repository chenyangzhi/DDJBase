all: build

fmt:
	gofmt -l -w -s src/

test:


build:dep
	go build -o bin/server src/main.go
	go build -o bin/client src/client.go

clean:
	rm -rf output

dep:fmt
