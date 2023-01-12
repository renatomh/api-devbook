package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// We can use this function to generate a random secret key
/*func init() {
	// Here, we'll generate the secret key
	key := make([]byte, 64)
	// Generating random data
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	// Showing the generated secret key
	fmt.Println(stringBase64)
}*/

func main() {
	// Loading environment vars
	config.Load()

	// Creating the router
	r := router.Generate()

	// Starting the server
	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
