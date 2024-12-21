.PHONY: test
test: ## Run unit tests with coverage
	@# NOTICE, the test output is using for coverage analytics, did not modify the std out
	@echo "cover package: ${UT_COVER_PACKAGES}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${UT_COVER_PACKAGES} -race -coverprofile cover.out -tags=\!integration ./...
	@go tool cover -func cover.out | tail -n 1 # print UT total coverage

.PHONY: test-raw
test-raw: ## Run unit tests without coverage(but all ut)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v -race -tags=\!integration ./...

.PHONY: integration-test
integration-test: ## Run integration tests
	@echo "cover package: ${IT_COVER_PACKAGES}"
	@if [ -d "./test/" ]; then \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${IT_COVER_PACKAGES} -coverprofile cover.out -race -tags=integration ./test/...; \
	fi

# Reference https://go.dev/testing/coverage/
.PHONY: show-coverage
show-coverage: ##  show coverage of UT and IT in specific packages
# No skip, run all tests
ifneq ($(skip), true)
	@# go version should be greater than @1.20
	@mkdir -p ${COVERAGE_PROFILING_DIR}
	@rm -rf ${COVERAGE_PROFILING_DIR}/*
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -cover -coverpkg=$(COVERAGE_PACKAGES) -tags=\!integration ${UT_COVER_PACKAGES} -args -test.gocoverdir=${COVERAGE_PROFILING_DIR} # Unit test
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -cover -coverpkg=$(COVERAGE_PACKAGES) -tags=integration ${IT_COVER_PACKAGES} -args -test.gocoverdir=${COVERAGE_PROFILING_DIR}	# Integration test
endif
	@go tool covdata func -i $(COVERAGE_PROFILING_DIR) -pkg $(COVERAGE_PACKAGES)| tail -n 1

