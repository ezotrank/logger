.PHONY: all deps format

all: deps format

deps:
	go get 'golang.org/x/tools/cmd/vet'
	go get

format:
	gofmt -w ./logger.go
	go tool vet -all=true ./logger.go

nuke:
	go clean -i

