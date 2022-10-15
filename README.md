# live-tests

This repository have unit tests, integration tests and contract tests. to be used as examples.

- Benchmarking - Does my code is faster than the prior implementation?
- Testing all possibilities - Fuzz test/ Mutation test
- Testing contracts - MockServer [Wiremock](https://wiremock.org)

## How to run it

```shell
go get ./...

go run cmd/server/main.go
```

## How to run tests

```shell
go test -v ./...
```

## How to check the tests coverage

```shell
go test -coverprofile cover.out ./... && go tool cover -html=cover.out
```
