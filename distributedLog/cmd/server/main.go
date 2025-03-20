package main

import (
	"log"

	"github.com/Shresth72/dslog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":6969")
	log.Fatal(srv.ListenAndServe())
}
