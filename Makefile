.PHONY: deps run-all

deps:
	go mod tidy

run-all:
	go run cmd/server/main.go & \
	go run cmd/client/main.go 