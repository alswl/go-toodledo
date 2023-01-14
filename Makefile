# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make              - default to 'build' target
#   make test         - run unit test
#   make build        - build local binary targets
#   make container    - build containers
#   make push         - push containers
#   make clean        - clean up targets
#
# The makefile is also responsible to populate project version information.

#
# Tweak the variables based on your project.
#
SHELL := /bin/bash
NOW_SHORT := $(shell date +%Y%m%d%H%M)

PROJECT := go-toodledo
# Target binaries. You can build multiple binaries for a single project.
TARGETS := toodledo tt

# Container registries.
REGISTRIES ?= ""

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

# This repo's root import path (under GOPATH).
ROOT := github.com/alswl/go-toodledo

# Project main package location (can be multiple ones).
CMD_DIR := ./cmd

# Project output directory.
OUTPUT_DIR := ./bin

# Build direcotory.
BUILD_DIR := ./build

# Git commit sha.
COMMIT := $(strip $(shell git rev-parse --short HEAD 2>/dev/null))
COMMIT := $(COMMIT)$(shell git diff-files --quiet || echo '-dirty')
COMMIT := $(if $(COMMIT),$(COMMIT),"Unknown")

# Current version of the project.
MAJOR_VERSION = 0
MINOR_VERSION = 1
PATCH_VERSION = 0
BUILD_VERSION = $(COMMIT)
GO_MOD_VERSION = $(shell cat go.mod | sha256sum | cut -c-6)
GOOS = $(shell go env GOOS)
GOARCH = ${shell go env GOARCH}
VERSION ?= v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)-$(BUILD_VERSION)

# TODO using k8s makefile style

all: download generate-code fmt test build

.PHONY: install-dev-tools
install-dev-tools:
	@bash ./hack/install-dev-tools.sh

.PHONY: download
download:
	go mod download

.PHONY: fmt
fmt:
	gofmt -w ./pkg/ ./test/ ./cmd/
	go fmt ./pkg/... ./cmd/...
	go vet ./pkg/... ./cmd/...

.PHONY: lint
lint:
	@echo "gofmt ensure"
	@test $$(gofmt -l ./pkg/ ./test/ ./cmd/ | wc -l) -eq 0

	@echo "ensure integration test with // +build integration tags"
	@test $$(find test -name '*_test.go' | wc -l) -eq $$(cat $$(find test -name '*_test.go') | grep -E '// ?go:build integration' | wc -l)

	golangci-lint run

.PHONY: generate-code
generate-code:
	@echo -n ''

.PHONY: generate-docs
generate-docs:
	rm -rf docs/toodledo/commands
	mkdir -p docs/commands
	go run ./cmd/gendocs/main.go

.PHONY: generate-code-enum
generate-code-enum:
	# go install golang.org/x/tools/cmd/stringer
	@echo generate stringer for enums
	# TODO using ls
	@(cd pkg/models/enums/tasks/priority; go generate)
	@(cd pkg/models/enums/tasks/status; go generate)
	@(cd pkg/models/enums/tasks/subtasksview; go generate)

.PHONY: generate-code-swagger
generate-code-swagger:
	@(cd pkg; rm client0/zz_generated_*.go;rm client0/*/zz_generated_*.go; rm models/zz_generated_*.go; swagger generate client -c client0 -f ../api/swagger.yaml -A toodledo --template-dir ../api/templates --allow-template-override -C ../api/config.yaml)

.PHONY: generate-code-mockery
generate-code-mockery:
	@echo generate mock of interfaces for testing
	@rm -rf test/mock
	@(cd pkg && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)
	@(cd cmd && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)

.PHONY: generate-code-wire
generate-code-wire:
	@echo generate wire
	@(cd cmd/toodledo/injector; $$GOPATH/bin/wire)

	@echo copy injector.go to itinjector.go for testing
	@mkdir -p test/suites/itinjector
	@cp cmd/toodledo/injector/injector.go test/suites/itinjector/itinjector.go
	@gsed -i 's/package injector/package itinjector/g' test/suites/itinjector/itinjector.go
	@gsed -i 's/TUISet/IntegrationTestTUISet/g' test/suites/itinjector/itinjector.go
	#@echo diff cmd/toodledo/injector/sets.go test/suites/itinjector/sets.go
	#diff cmd/toodledo/injector/sets.go test/suites/itinjector/sets.go || true
	@(cd test/suites/itinjector; $$GOPATH/bin/wire)

.PHONY: build
build:
	@for target in $(TARGETS); do                                                             \
	  GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) \
	    -ldflags "-s -w -X $(ROOT)/pkg/version.Version=$(VERSION)                             \
	    -X $(ROOT)/pkg/version.Commit=$(COMMIT)                                               \
	    -X $(ROOT)/pkg/version.Package=$(ROOT)"                                               \
	    $(CMD_DIR)/$${target};                                                                \
	  cp $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/$${target};                 \
	  cp $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/$${target}-$(VERSION);      \
	done

.PHONY: compress
compress:
	@for target in $(TARGETS); do   \
	  upx $(OUTPUT_DIR)/$${target}; \
	done

.PHONY: container
container:
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker build -t $${registry}$${image}:$(VERSION)                               \
	      --progress=plain                                                             \
	      -f $(BUILD_DIR)/$${target}/Dockerfile .;                                     \
	  done                                                                             \
	done

.PHONY: push
push: container
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker push $${registry}$${image}:$(VERSION);                                  \
	  done                                                                             \
	done

.PHONY: test
test: lint
	@go test -tags=\!integration ./...

.PHONY: integration-test
integration-test:
	@go test -tags=integration ./...

.PHONY: clean
clean:
	@rm -vrf ${OUTPUT_DIR}/*
