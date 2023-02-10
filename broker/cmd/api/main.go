package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 80

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %d.\n", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
