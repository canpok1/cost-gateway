package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/canpok1/code-gateway/internal/api"
	"github.com/canpok1/code-gateway/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := sql.Open("mysql", "service:password@tcp(db:3306)/cost-gateway?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	client := db.New(database)
	server := api.NewServer(client)

	r := http.NewServeMux()

	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
