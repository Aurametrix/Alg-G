ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif

REPO_NAME=template

PATH := $(GOPATH)/bin:$(PATH)
VERSION = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)

KIT_VERSION = "\
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.appName=${APP_NAME} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.version=${VERSION} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.branch=${BRANCH} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.revision=${REVISION} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.buildDate=${NOW} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.buildUser=${USER} \
	-X github.com/kolide/${REPO_NAME}/vendor/github.com/kolide/kit/version.goVersion=${GOVERSION}"

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux

	# If on macOS, set the shell to bash explicitly
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/bash
		CURRENT_PLATFORM = darwin
	endif

	# To populate version metadata, we use unix tools to get certain data
	GOVERSION = $(shell go version | awk '{print $$3}')
	NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
else
	CURRENT_PLATFORM = windows
	# To populate version metadata, we use windows tools to get the certain data
	GOVERSION_CMD = "(go version).Split()[2]"
	GOVERSION = $(shell powershell $(GOVERSION_CMD))
	NOW	= $(shell powershell Get-Date -format s)
endif

all: build
.PHONY: build
build: .pre-build .pre-template
	go build -i -o build/template -ldflags ${KIT_VERSION} ./cmd/template/

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux

.pre-template:
	$(eval APP_NAME = template)

xp: .pre-build .pre-template
	GOOS=darwin CGO_ENABLED=0 go build -i -o build/darwin/template -ldflags ${KIT_VERSION} ./cmd/template/
	GOOS=linux CGO_ENABLED=0 go build -i -o build/linux/template -ldflags ${KIT_VERSION} ./cmd/template/
	ln -f build/$(CURRENT_PLATFORM)/template build/template

test:
	go test -cover -race -v ./...
