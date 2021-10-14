package main

import (
	"log"

	"github.com/dmitrime/go-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
