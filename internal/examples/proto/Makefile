SHELL := /usr/bin/env bash

.PHONY: all
all: \
	clang-format \
	buf-lint \
	clang-format \
	buf-generate

include tools/buf/rules.mk
include tools/clang-format/rules.mk

build/protoc-gen-go: ../../../go.mod
	$(info [$@] building binary...)
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: build/protoc-gen-bq-json-schema
build/protoc-gen-bq-json-schema:
	$(info [$@] building binary...)
	@go build -o $@ ../../../protoc-gen-bq-json-schema

.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting proto schemas...)
	@$(buf) lint

.PHONY: clang-format
clang-format: $(clang_format)
	$(info [$@] formatting proto files...)
	@$(clang_format) \
		-i \
		--style='{BasedOnStyle: Google, ColumnLimit: 0, Language: Proto}' \
		$(shell find einride -type f -name '*.proto')

.PHONY: buf-generate
buf-generate: $(buf) build/protoc-gen-go build/protoc-gen-bq-json-schema
	$(info [$@] generating protobuf stubs...)
	@$(buf) generate \
		--template buf.gen.yaml \
		--path einride
