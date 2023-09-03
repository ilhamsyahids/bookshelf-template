package main

import (
	"log"
	"net/http"

	"github.com/ilhamsyahids/bookshelf-template/rest"
)

const addr = ":8080"

func main() {
	// init API
	api, err := rest.NewAPI(rest.APIConfig{})
	if err != nil {
		log.Fatalf("unable to initialize api due: %v", err)
	}

	// init server
	server := &http.Server{
		Addr:    addr,
		Handler: api.GetHandler(),
	}
	// run server
	log.Printf("server is listening on %v", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to run server due: %v", err)
	}
}
