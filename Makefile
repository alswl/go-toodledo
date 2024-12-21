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

# Tweak the variables based on your project.

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

## Current version of the project.
MAJOR_VERSION = 0
MINOR_VERSION = 1
PATCH_VERSION = 2
BUILD_VERSION = $(COMMIT)
GO_MOD_VERSION = $(shell cat go.mod | sha256sum | cut -c-6)
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
VERSION ?= v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)-$(BUILD_VERSION)

# Current version of the project.
VERSION_IN_FILE = $(shell cat VERSION)

UT_COVER_PACKAGES := $(shell go list ./pkg/... | grep -Ev 'pkg/clientsets|pkg/dal|pkg/models|pkg/version|pkg/injector')

.PHONY: all
all: fmt test build

include hack/makefile-go/_git.mk

.PHONY: install-dev-tools
install-dev-tools:
	echo ''

include hack/makefile-go/build.mk
include hack/makefile-go/test.mk
include hack/makefile-go/gen.mk
include hack/makefile-go/container.mk
include hack/makefile-go/version.mk
