package test

import (
	"database/sql"
	"live-tests/internal/infra/controller"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	sqlStatement := `
		CREATE TABLE IF NOT EXISTS clients (
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			email TEXT,
			points INTEGER
		);
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return db
}

func tearDownDb(db *sql.DB) {
	db.Exec("DROP TABLE clients")
	db.Close()
}

func TestCreateClientHandler(t *testing.T) {
	db := setupDb()
	defer tearDownDb(db)

	controller := controller.NewBaseHandler(db)

	t.Run("Should create a client", func(t *testing.T) {
		data := `{"name":"John Doe", "email": "j@j.com"}`
		reader := strings.NewReader(data)
		request, _ := http.NewRequest("POST", "/clients", reader)
		response := httptest.NewRecorder()
		controller.CreateClientHandler(response, request)
		if response.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.Code)
		}
	})
}
