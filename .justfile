# https://just.systems/man/en/

[private]
default:
    @just --list

# Run all automated code modifications.
checks: tidy vendor fmt

# Tidy dependencies.
tidy:
    go mod tidy

# Vendor dependencies.
vendor:
    go mod vendor

# Run all formatters.
fmt: fmt-go fmt-just

# Run the Go formatter.
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
fmt-just:
    just --unstable --fmt

# Check for uncommitted changes.
[private]
dirty:
    git diff --exit-code

# Run the linter.
lint:
    golangci-lint run

# Run the linter fixer.
lint-fix:
    golangci-lint run --fix

# Run the Go unit tests.
unit:
    go test -count=1 -cover ./...
