package main

import (
	"net/http"
	"test-go/config"
	"test-go/handlers"
	"test-go/migrations"
)

func main() {
	config.DbConfiguration()
	migrations.IndexMigration()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.IndexRouting(),
	}
	http.ListenAndServe(":8080", server.Handler)

}
