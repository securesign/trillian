.PHONY:
createtree-cross-platform: createtree-darwin-arm64 createtree-darwin-amd64 createtree-linux-amd64 createtree-linux-arm64 createtree-linux-ppc64le createtree-linux-s390x createtree-windows-amd64 ## Build all distributable (cross-platform) binaries

.PHONY:
updatetree-cross-platform: updatetree-darwin-arm64 updatetree-darwin-amd64 updatetree-linux-amd64 updatetree-linux-arm64 updatetree-linux-ppc64le updatetree-linux-s390x updatetree-windows-amd64 ## Build all distributable (cross-platform) binaries

.PHONY:	createtree-darwin-arm64
createtree-darwin-arm64: ## Build for mac M1
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o createtree-darwin-arm64 -trimpath ./cmd/createtree

.PHONY: createtree-darwin-amd64
createtree-darwin-amd64:  ## Build for Darwin (macOS)
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o createtree-darwin-amd64 -trimpath ./cmd/createtree

.PHONY: createtree-linux-amd64
createtree-linux-amd64: ## Build for Linux amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o createtree-linux-amd64 -trimpath ./cmd/createtree

.PHONY: createtree-linux-arm64
createtree-linux-arm64: ## Build for Linux arm64
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o createtree-linux-arm64 -trimpath ./cmd/createtree

.PHONY: createtree-linux-ppc64le
createtree-linux-ppc64le: ## Build for Linux ppc64le
	env CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -v -o createtree-linux-ppc64le -trimpath ./cmd/createtree

.PHONY: createtree-linux-s390x
createtree-linux-s390x:  ## Build for Linux s390x
	env CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -v -o createtree-linux-s390x -trimpath ./cmd/createtree

.PHONY: createtree-windows-amd64
createtree-windows-amd64: ## Build for Windows
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o createtree-windows-amd64.exe -trimpath ./cmd/createtree

.PHONY:	updatetree-darwin-arm64
updatetree-darwin-arm64: ## Build for mac M1
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o updatetree-darwin-arm64 -trimpath ./cmd/updatetree

.PHONY: updatetree-darwin-amd64
updatetree-darwin-amd64:  ## Build for Darwin (macOS)
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o updatetree-darwin-amd64 -trimpath ./cmd/updatetree

.PHONY: updatetree-linux-amd64
updatetree-linux-amd64: ## Build for Linux amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o updatetree-linux-amd64 -trimpath ./cmd/updatetree

.PHONY: updatetree-linux-arm64
updatetree-linux-arm64: ## Build for Linux arm64
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o updatetree-linux-arm64 -trimpath ./cmd/updatetree

.PHONY: updatetree-linux-ppc64le
updatetree-linux-ppc64le: ## Build for Linux ppc64le
	env CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -v -o updatetree-linux-ppc64le -trimpath ./cmd/updatetree

.PHONY: updatetree-linux-s390x
updatetree-linux-s390x:  ## Build for Linux s390x
	env CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -v -o updatetree-linux-s390x -trimpath ./cmd/updatetree

.PHONY: updatetree-windows-amd64
updatetree-windows-amd64: ## Build for Windows
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o updatetree-windows-amd64.exe -trimpath ./cmd/updatetree
