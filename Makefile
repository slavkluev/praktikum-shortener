BUILD_VERSION := v1.0.0
BUILD_DATE := $(shell date +'%Y/%m/%d')
BUILD_COMMIT := $(shell git rev-parse HEAD)

build:
	go build -ldflags "-X main.buildVersion=$(BUILD_VERSION) -X main.buildDate=$(BUILD_DATE) -X main.buildCommit=$(BUILD_COMMIT)" -o bin/shortener cmd/shortener/main.go
