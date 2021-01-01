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

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@go test -race -cover ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: buf-check-lint
buf-check-lint: $(buf)
	$(info [$@] linting protobuf schemas...)
	@$(buf) check lint

.PHONY: buf-generate
buf-generate: $(buf) $(protoc) $(protoc_gen_go)
	$(info [$@] generating protobuf stubs...)
	@$(buf) generate
