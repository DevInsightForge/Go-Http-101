# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Name of the executable
BINARY_NAME=gohttp101

# Path names
ENTRY_PATH=./cmd/main.go
OUT_PATH=./bin/

# Source files
SOURCES=$(wildcard *.go) $(wildcard */*.go)

all: clean build-cross

build: 
	$(GOBUILD) -o $(OUT_PATH)$(BINARY_NAME) $(ENTRY_PATH)

clean: 
	$(GOCLEAN)
	rm -rf $(OUT_PATH)

run: build
	$(OUT_PATH)$(BINARY_NAME)

run-dev:
	wgo run $(ENTRY_PATH)

build-cross: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)linux_amd64/$(BINARY_NAME) $(ENTRY_PATH)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)windows_amd64/$(BINARY_NAME).exe $(ENTRY_PATH)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUT_PATH)darwin_amd64/$(BINARY_NAME) $(ENTRY_PATH)
