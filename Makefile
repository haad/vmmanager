SHELL := /bin/bash
GIT_SHA=`git rev-parse --short HEAD || echo`

.DEFAULT_GOAL := all
.PHONY: all
all: ## build pipeline
all: mod inst gen seed build  #spell

.PHONY: ci
ci: ## CI build pipeline
ci: all lint test diff

.PHONY: release
br: all release

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -rf dist
	rm -f coverage.*
	rm -f '"$(shell go env GOCACHE)/../golangci-lint"'
	go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	go mod tidy

.PHONY: inst
inst: ## go install tools
	$(call print-target)

.PHONY: gen
gen: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: build
build:
	$(call print-target)
	mkdir -p build
	go build -ldflags "-X cmd.GitSHA=${GIT_SHA}" -o build/vmmanager .

.PHONY: release
## goreleaser build
release:
	$(call print-target)
	goreleaser build --rm-dist --single-target --snapshot

# .PHONY: spell
# spell: ## misspell
# 	$(call print-target)
# 	misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	golangci-lint run --fix

.PHONY: test
test: ## go test
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

.PHONY: seed
seed:
	$(call print-target)
	[ -f dist/seed.iso ] || hdiutil makehybrid -o dist/seed.iso -hfs -ov -joliet -iso -default-volume-name cidata dist/seed

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
