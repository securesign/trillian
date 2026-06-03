.PHONY:
createtree-cross-platform: createtree-darwin-arm64 createtree-darwin-amd64 createtree-windows-amd64 ## Build all distributable (cross-platform) binaries

# ---Darwin--- #
.PHONY: createtree-darwin-amd64
createtree-darwin-amd64:  ## Build for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=mod -o createtree-darwin-amd64 -trimpath ./cmd/createtree

.PHONY:	createtree-darwin-arm64
createtree-darwin-arm64: ## Build for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=mod -o createtree-darwin-arm64 -trimpath ./cmd/createtree

# ---Windows--- #
.PHONY: createtree-windows-amd64
createtree-windows-amd64: ## Build for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=mod -o createtree-windows-amd64.exe -trimpath ./cmd/createtree

.PHONY:
createtree-fips-cross-platform: createtree-fips-darwin-arm64 createtree-fips-darwin-amd64 createtree-fips-windows-amd64 ## Build all distributable FIPS (cross-platform) binaries

# ---Darwin--- #
.PHONY: createtree-fips-darwin-amd64
createtree-fips-darwin-amd64:  ## Build FIPS for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=mod -tags fips,no_openssl -o createtree-fips-darwin-amd64 -trimpath ./cmd/createtree

.PHONY:	createtree-fips-darwin-arm64
createtree-fips-darwin-arm64: ## Build FIPS for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=mod -tags fips,no_openssl -o createtree-fips-darwin-arm64 -trimpath ./cmd/createtree

# ---Windows--- #
.PHONY: createtree-fips-windows-amd64
createtree-fips-windows-amd64: ## Build FIPS for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=mod -tags fips,no_openssl -o createtree-fips-windows-amd64.exe -trimpath ./cmd/createtree

.PHONY:
updatetree-cross-platform: updatetree-darwin-arm64 updatetree-darwin-amd64 updatetree-windows-amd64 ## Build all distributable (cross-platform) binaries

# ---Darwin--- #
.PHONY:	updatetree-darwin-arm64
updatetree-darwin-arm64: ## Build for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=mod -o updatetree-darwin-arm64 -trimpath ./cmd/updatetree

.PHONY: updatetree-darwin-amd64
updatetree-darwin-amd64:  ## Build for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=mod -o updatetree-darwin-amd64 -trimpath ./cmd/updatetree

# ---Windows--- #
.PHONY: updatetree-windows-amd64
updatetree-windows-amd64: ## Build for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=mod -o updatetree-windows-amd64.exe -trimpath ./cmd/updatetree

.PHONY:
updatetree-fips-cross-platform: updatetree-fips-darwin-arm64 updatetree-fips-darwin-amd64 updatetree-fips-windows-amd64 ## Build all distributable FIPS (cross-platform) binaries

# ---Darwin--- #
.PHONY:	updatetree-fips-darwin-arm64
updatetree-fips-darwin-arm64: ## Build FIPS for mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=mod -tags fips,no_openssl -o updatetree-fips-darwin-arm64 -trimpath ./cmd/updatetree

.PHONY: updatetree-fips-darwin-amd64
updatetree-fips-darwin-amd64:  ## Build FIPS for Darwin (macOS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=mod -tags fips,no_openssl -o updatetree-fips-darwin-amd64 -trimpath ./cmd/updatetree

# ---Windows--- #
.PHONY: updatetree-fips-windows-amd64
updatetree-fips-windows-amd64: ## Build FIPS for Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=mod -tags fips,no_openssl -o updatetree-fips-windows-amd64.exe -trimpath ./cmd/updatetree
