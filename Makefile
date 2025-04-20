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
