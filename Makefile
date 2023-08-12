# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
.PHONY: clear
clear:
	rm -rf bin

## gen-docs: Swagger gen docs
.PHONY: gen-docs
gen-docs:
	swag init

## lint: Check code convention
.PHONY: lint
lint:
	golangci-lint run

## build: build the application
.PHONY: build
build:
	go build -o bin/main main.go

## run: run the application after built
.PHONY: run
run:
	./bin/main 

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

## start/dev: run the application
.PHONY: start/dev
start/dev: clear build run