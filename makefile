all: gen build run

gen: generate

generate: 
	go generate ./...

build:
	go build ./cmd/grimoireapi
run:
	./grimoireapi.exe