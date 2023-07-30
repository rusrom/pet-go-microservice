package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8090"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Statrting broker service on port %s\n", webPort)

	// define http server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
