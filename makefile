build-darwin:
	GOOS=darwin GOARCH=arm64 go build -o bin/preview-lite-arm64-darwin main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/preview-lite-amd64.exe main.go

build-all: build-darwin build-windows

.PHONY: build-darwin build-windows build-all