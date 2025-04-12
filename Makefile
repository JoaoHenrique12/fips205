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
	./golangci-lint-2.0.2-linux-amd64 run ./...

.PHONY: lint-fix
lint-fix:
	./golangci-lint-2.0.2-linux-amd64 run --fix
