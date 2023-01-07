package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Creating the router
	r := router.Generate()

	// Starting the server
	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

}
