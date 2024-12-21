.PHONY: generate-code-swagger
generate-code-swagger: ## generate code from swagger
	@echo generate swagger
	# TODO enable validation
	@(cd pkg; rm client/zz_generated_*.go; rm client/*/zz_generated_*.go; rm models/zz_generated_*.go; swagger generate client -f ../api/openapi.yaml -A $(PROJECT) -C ../api/config.yaml --skip-validation)
	# TODO waiting upstream merge
#	@(cd pkg; rm client/zz_generated_*.go; rm client/*/zz_generated_*.go; rm models/zz_generated_*.go; /Users/sigma-sre/dev/project/go-swagger/bin/swagger generate client -f ../api/openapi-fixed.json -A aone -C ../api/config.yaml --skip-validation)

.PHONY: generate-code-mockery
generate-code-mockery: ## Run generate generated unit test code
	# 如果遇到问题
	# Unexpected package creation during export data loading
	# https://github.com/vektra/mockery/pull/435#issuecomment-1134329306
	@echo "# generate mock of interfaces for testing"
	@rm -rf test/mock
	@mkdir -p test/mock
	@(cd . && mockery --all --keeptree --case=underscore --packageprefix=mock --output=./test/mock/)
	# mockery not support 1.18 generic now, temporarily drop zero size golang file
	# https://github.com/vektra/mockery/pull/456
	find test/mock -size 0 -exec rm {} \;


generate-code-enum: ## Generate enum String for models
	@echo generate stringer for enums
	# TODO using ls
	# @(cd pkg/models/enums/component/; go generate)


.PHONY: generate-manual
generate-manual: ## Generate develop docs
	@# go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	@# gomarkdoc --output MANUAL.md code.alipay.com/sigma-sre/hyper-common-utils

	@# if contains go files in .
	@if [ -n "$(shell ls *.go 2>/dev/null)" ]; then \
		echo "generate manual";                     \
		gomarkdoc --output MANUAL.md .;             \
	fi
	@if [ -d "cmd/gendoc/" ]; then \
  		echo "generate cmd doc";   \
  		mkdir -p docs;             \
		go run ./cmd/gendoc;       \
	fi

.PHONY: generate-gorm
generate-gorm:
	@if [ -d "cmd/dalgen/" ]; then \
		echo "generate gorm";      \
		go run $(ROOT)/cmd/dalgen; \
	fi

.PHONY: generate-code-wire
generate-code-wire: ## Generate wire injection code
	# wire
	@echo generate wire

	@for f in $(shell find . -name wire.go); do   \
		(cd $$(dirname $$f); $$GOPATH/bin/wire);  \
		echo "generate wire for $$(dirname $$f)"; \
	done

NIRVANA_BIN = $$HOME/local/bin/nirvana
.PHONY: generate-code-nv-openapiv3
generate-code-nv-openapiv3:  ## Generate OpenAPIV3
	@for target in $(TARGETS); do                                                                                       \
		if [ -f configs/$${target}/nirvana-it.yaml ]; then                                                              \
		  echo "# overwrite nirvana.yaml";                                                                              \
		  cp configs/$${target}/nirvana-it.yaml nirvana.yaml;                                                           \
        fi;                                                                                                             \
		if [ -d pkg/web/$${target}/apis ]; then                                                                         \
			  echo "# generate pkg/web/$${target}/apis";                                                                \
			$(NIRVANA_BIN) api pkg/web/$${target}/apis --serve= --output=/tmp --escape-class-name-symbol --open-api-v3; \
			mv /tmp/api.v1.json api/openapi/$${target}-api.json;                                                        \
		fi;                                                                                                             \
	done

