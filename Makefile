# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod

# Name of the executable
BINARY_NAME=gohttp101

# Path names
ENTRY_PATH=./cmd/api/main.go
OUT_PATH=./bin/

# Source files
SOURCES=$(wildcard *.go) $(wildcard */*.go)

all: clean build-cross

deps:
	$(GOMOD) download
	$(GOMOD) verify

build: 
	$(GOBUILD) -o $(OUT_PATH)$(BINARY_NAME) $(ENTRY_PATH)

clean: 
	$(GOCLEAN)
	rm -rf $(OUT_PATH)

run: 
	$(OUT_PATH)$(BINARY_NAME)

dev:
	wgo run $(ENTRY_PATH)

build-cross: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)linux_amd64/$(BINARY_NAME) $(ENTRY_PATH)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)windows_amd64/$(BINARY_NAME).exe $(ENTRY_PATH)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)darwin_amd64/$(BINARY_NAME) $(ENTRY_PATH)
