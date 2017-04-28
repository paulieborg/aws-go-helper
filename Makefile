.PHONY: get build run

BUILD_DIR=./build

BINARY_FILE:=$(BUILD_DIR)/cloudformation

get:
	go get -u github.com/aws/aws-sdk-go

build: get
	@echo "+++ building for $(GOOS)-$(GOARCH) ..."
	go build -v -o $(BINARY_FILE)
	chmod 777 $(BUILD_DIR) && chmod +x $(BINARY_FILE)

run:
	build/cloudformation -n TestStack
