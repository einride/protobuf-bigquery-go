SHELL := /bin/bash

.PHONY: all
all: \
	commitlint \
	examples-proto \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	git-verify-nodiff

include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/semantic-release/rules.mk

.PHONY: clean
clean:
	$(info [$@] removing build files...)
	@rm -rf build

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@mkdir -p build/coverage
	@go test -short -race -coverprofile=build/coverage/$@.txt -covermode=atomic ./...

.PHONY: go-integration-test
go-integration-test:
	$(info [$@] running Go tests (including integration tests)...)
	@mkdir -p build/coverage
	@go test -race -cover -coverprofile=build/coverage/$@.txt -covermode=atomic ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: examples-proto
examples-proto:
	make -C internal/examples/proto
