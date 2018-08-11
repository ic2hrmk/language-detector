package main

import (
	"log"
	"flag"
	"net/http"

	"lang-detector/rest"
)

func main() {
	//
	// Reading configuration
	//
	port := flag.String("port", ":8080", "port to run on")
	flag.Parse()

	log.Printf("[detect-language] Server is started on port %s\n", *port)

	//
	// Init. web-server
	//
	rest.Init()

	//
	// Listen to connections
	//

	log.Fatal(http.ListenAndServe(":8080", nil))
}


