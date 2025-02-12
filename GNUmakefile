TEST?=$$(go list -e ./... | grep -v github.com/hashicorp/terraform-provider-google-beta/scripts)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=google
DIR_NAME=google-beta

default: build

build: lint
	go install

# Compile the analyzer
analyze:
	go build -o ./scripts/package-needs-unit-testing/pnut ./scripts/package-needs-unit-testing/main.go

test: lint testnolint

# Used in CI to prevent lint failures from being interpreted as test failures
testnolint: analyze
	@tmpfile=$$(mktemp -t testnolint.XXXXXX); \
	for pkg in $$(go list -e $(TEST)); do \
		dir=$$(go list -f '{{.Dir}}' "$$pkg"); \
		if [ -d "$$dir" ]; then \
			( \
				./scripts/package-needs-unit-testing/pnut "$$dir" | { \
					read line; \
					case "$$line" in \
						TESTABLE:*) \
							package_path=$$(echo "$$line" | cut -d':' -f2); \
							echo "TESTABLE:$$package_path" >> "$$tmpfile"; \
							;; \
						SKIPPED:*) \
							package_path=$$(echo "$$line" | cut -d':' -f2); \
							echo "Skipping tests for $$package_path (all tests use acctest.VcrTest)"; \
							;; \
					esac; \
				} \
			) & \
		fi; \
	done; \
	wait; \
	packages="$$(grep "^TESTABLE:" "$$tmpfile" | cut -d':' -f2 | tr '\n' ' ')"; \
	if [ -n "$$packages" ]; then \
		go test -timeout=30s $$packages; \
	else \
		echo "No packages to test."; \
	fi; \
	rm -f "$$tmpfile"

testacc: lint
	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 go test $(TEST) -v $(TESTARGS) -timeout 240m -ldflags="-X=github.com/hashicorp/terraform-provider-google-beta/version.ProviderVersion=acc"

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(DIR_NAME)

# Currently required by tf-deploy compile
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

vet:
	go vet

lint: fmtcheck vet

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(DIR_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

docscheck:
	@sh -c "'$(CURDIR)/scripts/docscheck.sh'"

.PHONY: build test testnolint testacc fmt fmtcheck vet lint  errcheck test-compile website website-test docscheck analyze