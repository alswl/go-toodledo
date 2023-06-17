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

# Build directory.
BUILD_DIR := ./build

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

# Git commit sha.
COMMIT := $(strip $(shell git rev-parse --short HEAD 2>/dev/null))
COMMIT := $(COMMIT)$(shell [[ -z $$(git status -s) ]] || echo '-dirty')
COMMIT := $(if $(COMMIT),$(COMMIT), $${COMMIT})
COMMIT := $(if $(COMMIT),$(COMMIT),"Unknown")

# Current version of the project.
MAJOR_VERSION = 0
MINOR_VERSION = 1
PATCH_VERSION = 0
BUILD_VERSION = $(COMMIT)
GO_MOD_VERSION = $(shell cat go.mod | sha256sum | cut -c-6)
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
VERSION ?= v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)-$(BUILD_VERSION)

UT_COVER_PACKAGES := $(shell go list ./pkg/... |grep -Ev 'pkg/clientsets|pkg/dal|pkg/models|pkg/version|pkg/injector')

.PHONY: all
all: fmt test build

.PHONY: check-git-status
check-git-status: ## Check git status
	@test -z "$$(git status --porcelain)" || (echo "Git status is not clean, please commit or stash your changes first" && exit 1)

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: install-dev-tools
install-dev-tools: ## Install dev tools
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	go install golang.org/x/tools/cmd/stringer
	bash ./hack/install-dev-tools.sh

##@ Build

.PHONY: download
download: ## Run go mod download
	go mod download

.PHONY: fmt
fmt: ## Run format code
	gofmt -w ./pkg/ ./test/ ./cmd/
	go fmt ./pkg/... ./cmd/...
	go vet ./pkg/... ./cmd/...
	goimports -w ./pkg/ ./test/ ./cmd/
	golangci-lint run --fix

.PHONY: lint
lint: ## Run lint
	@echo "# gofmt"
	@test $$(gofmt -l . | wc -l) -eq 0

	@echo "# ensure integration test with // +build integration"
	@test $$(find test -name '*_test.go' | wc -l) -eq $$(cat $$(find test -name '*_test.go') | grep -E '// ?go:build integration' | wc -l)

	go mod tidy
	golangci-lint run --timeout 5m
	gofmt -w .

.PHONY: generate-code
generate-code: ## deprecated , using others generate-code-x
	@echo -n ''

.PHONY: generate-code-swagger
generate-code-swagger: ## generate code from swagger
	@echo generate swagger
	@(cd pkg; rm client0/zz_generated_*.go;rm client0/*/zz_generated_*.go; rm models/zz_generated_*.go; swagger generate client -c client0 -f ../api/swagger.yaml -A toodledo --template-dir ../api/templates --allow-template-override -C ../api/config.yaml)

.PHONY: generate-code-mockery
generate-code-mockery: ## Run generate generated unit test code
	# 如果遇到问题
	# Unexpected package creation during export data loading
	# https://github.com/vektra/mockery/pull/435#issuecomment-1134329306
	@echo "# generate mock of interfaces for testing"
	@rm -rf test/mock
	@mkdir -p test/mock
	#@(cd pkg && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)
	#@(cd cmd && mockery --all --keeptree --case=underscore --packageprefix=mock --output=../test/mock)
	@(cd . && mockery --all --keeptree --case=underscore --packageprefix=mock --output=./test/mock/)
	# mockery not support 1.18 generic now, temporarily drop zero size golang file
	# https://github.com/vektra/mockery/pull/456
	find test/mock -size 0 -exec rm {} \;

.PHONY: generate-code-wire
generate-code-wire: ## Generate wire injection code
	# wire
	@echo generate wire
	@(cd cmd/toodledo/injector; $$GOPATH/bin/wire)

	@echo copy injector.go to itinjector.go for testing
	@mkdir -p test/suites/itinjector
	@cp cmd/toodledo/injector/injector.go test/suites/itinjector/itinjector.go
	@gsed -i 's/package injector/package itinjector/g' test/suites/itinjector/itinjector.go
	@gsed -i 's/TUISet/IntegrationTestTUISet/g' test/suites/itinjector/itinjector.go
	@(cd test/suites/itinjector; $$GOPATH/bin/wire)

.PHONY: generate-code-enum
generate-code-enum: ## Generate enum String for models
	# go install golang.org/x/tools/cmd/stringer
	@echo generate stringer for enums
	# TODO using ls
	@(cd pkg/models/enums/tasks/priority; go generate)
	@(cd pkg/models/enums/tasks/status; go generate)
	@(cd pkg/models/enums/tasks/subtasksview; go generate)

.PHONY: generate-manual
generate-manual: ## Generate develop docs
	# go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc --output GO_DOC.md ./...

	# TODO for loop
	rm -rf docs/toodledo/commands
	mkdir -p docs/commands
	go run ./cmd/gendocs/

.PHONY: build
MINIMAL_BUILD = 0
build: ## Build
	@for target in $(TARGETS); do                                                             \
	  GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) \
	    -ldflags "-s -w -X $(ROOT)/pkg/version.Version=$(VERSION)                             \
	    -X $(ROOT)/pkg/version.Commit=$(COMMIT)                                               \
	    -X $(ROOT)/pkg/version.Package=$(ROOT)"                                               \
	    $(CMD_DIR)/$${target};                                                                \
	  if [ $(MINIMAL_BUILD) -eq 1 ]; then                                                                                  \
	    upx $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH);                                                                    \
	  fi;                                                                                                                  \
	  cp $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/$${target}-$(VERSION)-$(GOOS)-$(GOARCH);                 \
	  cp $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/$${target}-$(VERSION);                                   \
	  cp $(OUTPUT_DIR)/$${target}-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/$${target};                                              \
	done

.PHONY: compress
compress: ## Compress by upx
	@for target in $(TARGETS); do   \
	  upx $(OUTPUT_DIR)/$${target}; \
	done


.PHONY: container-build-env
newBuildEnvVersion=$(GO_MOD_VERSION)-$(shell tar cf - build/$(PROJECT) | sha256sum | cut -c1-6)
container-build-env: #check-git-status ## Build container build env
	@for registry in $(REGISTRIES); do                                                          \
	  image=$(IMAGE_PREFIX)$(PROJECT)$(IMAGE_SUFFIX);                                           \
	    DOCKER_BUILDKIT=1 docker build -t $${registry}$${image}-build-env:$(newBuildEnvVersion) \
	      --build-arg ROOT=$(ROOT) --build-arg TARGET=$${target}                                \
	      --build-arg CMD_DIR=$(CMD_DIR)                                                        \
	      --build-arg VERSION=$(GO_MOD_VERSION)                                                 \
	      --build-arg COMMIT=$(COMMIT)                                                          \
	    --progress=plain                                                                        \
	    -f $(BUILD_DIR)/$(PROJECT)/Dockerfile .;                                                \
	done

.PHONY: container
container: ## Build containers
	# COMMIT pass to container, because no git repo in container
	@for target in $(TARGETS); do                                \
	  for registry in $(REGISTRIES); do                          \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);          \
	    docker build -t $${registry}$${image}:$(VERSION)         \
	      --build-arg COMMIT=$(COMMIT)                   \
	      --progress=plain                                       \
	      -f $(BUILD_DIR)/$${target}/Dockerfile .;               \
	  done                                                       \
	done

.PHONY: push-container-build-env
newBuildEnvVersion=$(GO_MOD_VERSION)-$(shell tar cf - build/$(PROJECT) | sha256sum | cut -c1-6)
push-container-build-env: ## Push containers build env images to reigstry
	@for registry in $(REGISTRIES); do                                     \
	  image=$(IMAGE_PREFIX)$(PROJECT)$(IMAGE_SUFFIX);                      \
	    docker push $${registry}$${image}-build-env:$(newBuildEnvVersion); \
	done

.PHONY: bump-build-env-container
currentBuildEnvVersion=$(shell head -n 1 build/$(PROJECT)/VERSION)
newBuildEnvVersion=$(GO_MOD_VERSION)-$(shell tar cf - build/$(PROJECT) | sha256sum | cut -c1-6)
bump-build-env-container: ## Bump build env container for all containers and aci
	echo "$(newBuildEnvVersion)" > build/$(PROJECT)/VERSION
	gsed -i "s/build-env:$(currentBuildEnvVersion)/build-env:$(newBuildEnvVersion)/g" .aci.yml
	find build -name "Dockerfile" -exec gsed -i "s/build-env:$(currentBuildEnvVersion)/build-env:$(newBuildEnvVersion)/g" {} \;
	@echo "# PLEASE using 'git commit -a' to commit image version changes"


.PHONY: push
push: ## Push containers images to reigstry
	@for target in $(TARGETS); do                       \
	  for registry in $(REGISTRIES); do                 \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX); \
	    docker push $${registry}$${image}:$(VERSION);   \
	  done                                              \
	done

.PHONY: test
test: ## Run unit tests
	@# NOTICE, the test output is using for coverage analytics, did not modify the std out
	@echo "cover package: ${UT_COVER_PACKAGES}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${UT_COVER_PACKAGES} -race -coverprofile cover.out -tags=\!integration ./...

.PHONY: integration-test
integration-test: ## Run integration tests
	@echo "cover package: ${IT_COVER_PACKAGES}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${IT_COVER_PACKAGES} -coverprofile cover.out -race -tags=integration ./...

.PHONY: clean
clean: ## Clean temp files
	@rm -vrf ${OUTPUT_DIR}/*
