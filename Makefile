
local-vault:
	vault server -dev

setup-test-data:
	./helpers/setup.sh

.PHONY: install-osx
install-osx:
	cp ./bin/vlt/vlt /usr/local/bin/vlt

.PHONY: dev
dev: ## Build for the current development version
	@echo "==> Building vlt..."
	@mkdir -p ./bin
	@CGO_ENABLED=0 go build -o ./bin/vlt ./cmd/vlt
	@rm -f $(GOPATH)/bin/vlt
	@cp ./bin/vlt/vlt $(GOPATH)/bin/vlt
	@echo "==> Done"

.PHONY: build
build:
	go build -ldflags "-X main.version=`git tag --sort=-version:refname | head -n 1`" -o bin/vlt ./cmd/vlt

.PHONY: run
run:
	./bin/vlt

.PHONY: test
test:
	go test ./...