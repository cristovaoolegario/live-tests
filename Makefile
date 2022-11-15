setup:
	$(MAKE) create-db && $(MAKE) install-deps

create-db:
	sqlite3 db.sqlite ".read ./tools/create-db.sql"

install-deps:
	go mod tidy

tests:
	go clean -testcache
	go test -v ./...

test-coverage:
	go test -coverprofile cover.out ./... && go tool cover -html=cover.out

run:
	go run cmd/server/main.go