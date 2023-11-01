# Go project variables

MAIN_FILE := main.go
BINARY_NAME := kvdb
TEST_DIR := tests

#--------------------------------------make clean--------------------------------------

# Detect the current operating system
ifeq ($(OS),Windows_NT)
    CLEAN_COMMAND := clean-windows
else
    CLEAN_COMMAND := clean-linux
endif

# Clean the project (remove the compiled binary)
.PHONY: clean
clean: $(CLEAN_COMMAND)

# Clean for Linux
.PHONY: clean-linux
clean-linux:
	rm -f $(BINARY_NAME)

# Clean for Windows
.PHONY: clean-windows
clean-windows:
	del $(BINARY_NAME).exe

#--------------------------------------make install--------------------------------------

# Install project dependencies
.PHONY: install
install:
	go mod download

# Run tests
.PHONY: test
test:
	go test -v ./$(TEST_DIR)

#--------------------------------------make build--------------------------------------

# Detect the current operating system
ifeq ($(OS),Windows_NT)
    BUILD_TARGET := build-windows
else
    BUILD_TARGET := build-linux
endif

# Default target: build and run the project
.PHONY: build
build: $(BUILD_TARGET)

# Build the project for a Linux system
.PHONY: build-linux
build-linux:
	GOOS=linux 
	GOARCH=amd64 
	go build -o $(BINARY_NAME)

# Build the project for a Windows system
.PHONY: build-windows
build-windows:
	set GOOS=windows 
	set GOARCH=amd64 
	go build -o $(BINARY_NAME).exe

#--------------------------------------make run--------------------------------------

# Detect the current operating system
ifeq ($(OS),Windows_NT)
    RUN_TARGET := run-windows
else
    RUN_TARGET := run-linux
endif

# Default target: run the project
.PHONY: run
run: $(RUN_TARGET)

# Run the project for a Linux system
.PHONY: run-linux
run-linux:
	./$(BINARY_NAME) set mykey myvalue

# Run the project for a Windows system
.PHONY: run-windows
run-windows:
	./$(BINARY_NAME) set mykey myvalue

# Shortcut for installing dependencies, building, and running
.PHONY: all
all: clean install build test run

