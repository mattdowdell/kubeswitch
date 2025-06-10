# https://just.systems/man/en/

[private]
default:
    @just --list

# Run the checks, lint and unit recipes.
all: checks lint unit

# Run all automated code modifications.
checks: tidy vendor fmt

# Tidy dependencies.
[group('dependencies')]
tidy:
    go mod tidy

# Vendor dependencies.
[group('dependencies')]
vendor:
    go mod vendor

# Run all formatters.
[group('formatters')]
fmt: fmt-go fmt-just

# Run the Go formatter.
[group('formatters')]
fmt-go:
    gofumpt -l -w .
    gci write \
        --skip-vendor \
        --skip-generated \
        -s standard \
        -s default \
        -s localmodule \
        .

# Run the Justfile formatter.
[group('formatters')]
fmt-just:
    just --unstable --fmt

# Check for uncommitted changes.
[private]
dirty:
    git diff --exit-code

# Run the linter.
[group('linters')]
lint:
    golangci-lint run

# Run the linter fixer.
[group('linters')]
lint-fix:
    golangci-lint run --fix

# Run the Go unit tests.
[group('tests')]
unit:
    go test -count=1 -cover -coverprofile=unit.out ./...
    @echo "Total coverage: `go tool cover -func=unit.out | tail -n 1 | awk '{print $3}'`"
    go tool cover -html unit.out -o unit.html
