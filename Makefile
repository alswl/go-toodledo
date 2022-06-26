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
GOOS = darwin
GOARCH = amd64
VERSION ?= v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)-$(BUILD_VERSION)



.PHONY: fmt build container test integration-test push clean lint download generate-code compress

all: download generate-code fmt test build

download:
	go mod download

fmt:
	gofmt -w ./pkg/ ./test/ ./cmd/
	go fmt ./pkg/... ./cmd/...
	go vet ./pkg/... ./cmd/...

lint:
	@echo "gofmt ensure"
	@test $$(gofmt -l ./pkg/ ./test/ ./cmd/ | wc -l) -eq 0

	@echo "ensure integration test with // +build integration tags"
	@test $$(find test -name '*_test.go' | wc -l) -eq $$(cat $$(find test -name '*_test.go') | grep -E '// ?\+build integration' | wc -l)

	# FIXME enable it
	# golangci-lint run

generate-code:
	@echo -n ''

generate-code-enum:
	# go install golang.org/x/tools/cmd/stringer
	@echo generate stringer for enums
	# TODO using ls
	@(cd pkg/models/enums/tasks/priority; go generate)
	@(cd pkg/models/enums/tasks/status; go generate)
	@(cd pkg/models/enums/tasks/subtasksview; go generate)

generate-code-swagger:
	@(cd pkg; rm client/zz_generated_*.go;rm client/*/zz_generated_*.go; rm models/zz_generated_*.go; swagger generate client -f ../api/swagger.yaml -A toodledo --template-dir ../api/templates --allow-template-override -C ../api/config.yaml)

generate-code-mockery:
	@echo generate mock of interfaces for testing
	@rm -rf test/mock
	@(cd pkg && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)
	@(cd cmd && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)

generate-code-wire:
	@echo generate wire
	@(cd cmd/toodledo/injector; $$GOPATH/bin/wire)

	@echo copy injector.go to itinjector.go for testing
	@mkdir -p test/suites/itinjector
	@cp cmd/toodledo/injector/injector.go test/suites/itinjector/itinjector.go
	@gsed -i 's/package injector/package itinjector/g' test/suites/itinjector/itinjector.go
	@gsed -i 's/SuperSet/IntegrationTestSet/g' test/suites/itinjector/itinjector.go
	#@echo diff cmd/toodledo/injector/sets.go test/suites/itinjector/sets.go
	#diff cmd/toodledo/injector/sets.go test/suites/itinjector/sets.go || true
	@(cd test/suites/itinjector; $$GOPATH/bin/wire)

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

compress:
	@for target in $(TARGETS); do   \
	  upx $(OUTPUT_DIR)/$${target}; \
	done

container:
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker build -t $${registry}$${image}:$(VERSION)                               \
	      --progress=plain                                                             \
	      -f $(BUILD_DIR)/$${target}/Dockerfile .;                                     \
	  done                                                                             \
	done

push: container
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker push $${registry}$${image}:$(VERSION);                                  \
	  done                                                                             \
	done

test: lint
	@go test -tags=\!integration ./...

integration-test:
	@go test -tags=integration ./...

clean:
	@rm -vrf ${OUTPUT_DIR}/*
