.PHONY: glide build run

BUILD_DIR=./bin

BINARY_FILE:=$(BUILD_DIR)/cloudformation

glide:
	glide install

build: glide
	@echo "+++ building for $(GOOS)-$(GOARCH) ..."
	go build -v -o $(BINARY_FILE)
	chmod 755 $(BUILD_DIR) && chmod +x $(BINARY_FILE)

run:
	#Use defaults -n MyTestStack -t templates/template.yml -p templates/params.json
	bin/cloudformation -n MyTestStack -a $(ACTION)

unit_test: glide
	go test $(shell go list ./... | grep -v /vendor/)

integration_test:
	scripts/test.sh

test: unit_test integration_test