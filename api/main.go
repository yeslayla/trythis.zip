package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yeslayla/trythis.zip/api/app"
)

func main() {
	app := &app.App{}

	server := &http.Server{
		Handler: app.GetRouter(),
		Addr:    "0.0.0.0:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
