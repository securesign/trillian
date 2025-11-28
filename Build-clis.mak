FIPS_MODULE ?= latest

.PHONY:
createtree-cross-platform: createtree-darwin-arm64 createtree-darwin-amd64 createtree-windows-amd64 ## Build all distributable (cross-platform) binaries

# ---Darwin--- #
.PHONY: createtree-darwin-amd64
createtree-darwin-amd64:  ## Build for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o createtree-darwin-amd64 -trimpath ./cmd/createtree

.PHONY:	createtree-darwin-arm64
createtree-darwin-arm64: ## Build for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o createtree-darwin-arm64 -trimpath ./cmd/createtree

# ---Windows--- #
.PHONY: createtree-windows-amd64
createtree-windows-amd64: ## Build for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o createtree-windows-amd64.exe -trimpath ./cmd/createtree

.PHONY:
updatetree-cross-platform: updatetree-darwin-arm64 updatetree-darwin-amd64 updatetree-windows-amd64 ## Build all distributable (cross-platform) binaries

# ---Darwin--- #
.PHONY:	updatetree-darwin-arm64
updatetree-darwin-arm64: ## Build for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o updatetree-darwin-arm64 -trimpath ./cmd/updatetree

.PHONY: updatetree-darwin-amd64
updatetree-darwin-amd64:  ## Build for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o updatetree-darwin-amd64 -trimpath ./cmd/updatetree

# ---Windows--- #
.PHONY: updatetree-windows-amd64
updatetree-windows-amd64: ## Build for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOFIPS140=$(FIPS_MODULE) go build -mod=mod -o updatetree-windows-amd64.exe -trimpath ./cmd/updatetree
