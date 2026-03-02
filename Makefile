.PHONY: all
all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64

.PHONY: build-dir
build-dir:
	mkdir -p build

.PHONY: build-darwin-amd64
build-darwin-amd64: docs build-dir
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/fibo_darwin_amd64 ./cmd/api

.PHONY: build-darwin-amd64
build-darwin-arm64: docs build-dir
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/fibo_darwin_arm64 ./cmd/api

.PHONY: build-linux-amd64
build-linux-amd64: docs build-dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/fibo_linux_amd64 ./cmd/api

.PHONY: build-linux-arm64
build-linux-arm64: docs build-dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/fibo_linux_arm64 ./cmd/api

.PHONY: build-windows-amd64
build-windows-amd64: docs build-dir
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/fibo_windows_amd64.exe ./cmd/api

.PHONY: build-docker-image
build-docker-image: docs
	docker buildx build --platform linux/amd64,linux/arm64 -f ./docker/Dockerfile -t majabojarska/fibo .

.PHONY: clean
clean:
	rm -rf build

.PHONY: docs
docs:
	swag fmt
	swag init --dir internal/routes/ --parseInternal --generalInfo router.go

.PHONY: test
test:
	FIBO_METRICS_ENABLED=false FIBO_DOCS_ENABLED=false go test -v ./...
