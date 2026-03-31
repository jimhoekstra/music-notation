alias cov := test-coverage

# Run the tests
test:
    go test ./...

# Run the tests with coverage
test-coverage:
    go test ./... -cover