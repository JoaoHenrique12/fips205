.PHONY: install-ci-tools
install-ci-tools:
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@if ! command -v golangci-lint >/dev/null || [ "$$(golangci-lint --version | grep -o 'v[0-9]\+\.[0-9]\+\.[0-9]\+')" != "v2.1.2" ]; then \
		echo "Installing golangci-lint v2.1.2..."; \
		curl -sSL https://github.com/golangci/golangci-lint/releases/download/v2.1.2/golangci-lint-2.1.2-linux-amd64.tar.gz | tar -xz; \
		sudo mv golangci-lint-2.1.2-linux-amd64/golangci-lint /usr/local/bin/; \
		rm -rf golangci-lint-2.1.2-linux-amd64
	else \
		echo "golangci-lint v2.1.2 already installed."; \
	fi

.PHONY: coverage
coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: coverage-inspect-html
coverage-inspect-html: coverage
	go tool cover -html=coverage.out -o coverage.html

.PHONY: coverage-inspect-text
coverage-inspect-text: coverage
	go tool cover -func=coverage.out

.PHONY: format
format:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: top5-cyclo
top5-cyclo:
	gocyclo -top 5 -ignore test_* ./
