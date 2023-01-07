package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Loading environment vars
	config.Load()

	// Creating the router
	r := router.Generate()

	// Starting the server
	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
