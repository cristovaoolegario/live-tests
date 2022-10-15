package main

import (
	"database/sql"
	"live-tests/internal/infra/controller"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	controllers := controller.NewBaseHandler(db)

	http.HandleFunc("/clients", controllers.CreateClientHandler)
	http.ListenAndServe(":5000", nil)

}
