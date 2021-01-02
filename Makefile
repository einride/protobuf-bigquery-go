SHELL := /bin/bash

.PHONY: all
all: \
	commitlint \
	buf-check-lint \
	buf-generate \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	git-verify-nodiff

include tools/buf/rules.mk
include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/protoc-gen-go/rules.mk
include tools/protoc/rules.mk
include tools/semantic-release/rules.mk

.PHONY: clean
clean:
	$(info [$@] removing build files...)
	@rm -rf build

.PHONY: internal/examples/proto/api-common-protos
internal/examples/proto/api-common-protos:
	@git submodule update --init --recursive $@

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

.PHONY: buf-check-lint
buf-check-lint: $(buf) internal/examples/proto/api-common-protos
	$(info [$@] linting protobuf schemas...)
	@$(buf) check lint

.PHONY: buf-generate
buf-generate: $(buf) $(protoc) $(protoc_gen_go) internal/examples/proto/api-common-protos
	$(info [$@] generating protobuf stubs...)
	@rm -rf internal/examples/proto/gen
	@$(buf) generate --path internal/examples/proto/src/einride
