.PHONY: all gxphere test fmt clean 

GOBIN = /usr/local/bin
GOOS ?= $(shell go env GOOS)
GORUN = go run

gxphere: ## Build the project
	mkdir -p $(GOBIN)
	$(GORUN) build/ci.go install -o $(GOBIN)/gxphere
	@echo "gxphere installed in $(GOBIN)."
	@export PATH=$(GOBIN):$$PATH
	@echo "PATH temporarily updated. Run 'export PATH=\$$PATH:$(GOBIN)' to persist."
	
all: gxphere

test: ## Run all tests globally
	go test ./...

# lint: ## Run linters
# 	$(GORUN) build/ci.go lint

# archive: ## Archive build artifacts
# 	$(GORUN) build/ci.go archive

# debsrc: ## Generate Debian source package
# 	$(GORUN) build/ci.go debsrc

# purge: ## Purge old build artifacts
# 	$(GORUN) build/ci.go purge

fmt: ## Format code
	gofmt -s -w $(shell find . -name "*.go")

clean: ## Clean up cache and binaries
	go clean -cache
	rm -rf $(GOBIN)