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

# Target binaries. You can build multiple binaries for a single project.
TARGETS := toodledo

# Container registries.
REGISTRIES ?= ""

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

# This repo's root import path (under GOPATH).
ROOT := alswl/go-toodledo

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
VERSION ?= v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)-$(BUILD_VERSION)-$(NOW_SHORT)

#
# Define all targets. At least the following commands are required:
#

.PHONY: build container push test integration clean swagger

build:
	@for target in $(TARGETS); do                                                      \
	  go build -v -o $(OUTPUT_DIR)/$${target}                                          \
	    -ldflags "-s -w -X $(ROOT)/pkg/version.Version=$(VERSION)                      \
	    -X $(ROOT)/pkg/version.Commit=$(COMMIT)                                        \
	    -X $(ROOT)/pkg/version.Package=$(ROOT)"                                        \
	    $(CMD_DIR)/$${target};                                                         \
	done
	

swagger:
	(cd pkg; swagger generate client -f ../api/swagger.yaml -A toodledo --template-dir ../api/templates)

container:
	@for target in $(TARGETS); do                                                      \
	  for registry in $(REGISTRIES); do                                                \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                \
	    docker build -t $${registry}$${image}:$(VERSION)                               \
	      --build-arg ROOT=$(ROOT) --build-arg TARGET=$${target}                       \
	      --build-arg CMD_DIR=$(CMD_DIR)                                               \
	      --build-arg VERSION=$(VERSION)                                               \
	      --build-arg COMMIT=$(COMMIT)                                                 \
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

test:
	@go test -tags=\!integration ./...

integration:
	@go test -tags=integration ./...

clean:
	@rm -vrf ${OUTPUT_DIR}/*
