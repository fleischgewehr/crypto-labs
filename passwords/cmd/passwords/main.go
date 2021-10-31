package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/server"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/shutdown"
)

func main() {
	http.HandleFunc("/", Handler)

	app, err := app.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.Get().WithAddr(":8080")

	go func() {
		log.Println("server started at localhost:8080")
		if err := srv.Start(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	shutdown.Exit(func() {
		if err := srv.Close(); err != nil {
			log.Fatal(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			log.Fatal(err.Error())
		}
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yo")
}
