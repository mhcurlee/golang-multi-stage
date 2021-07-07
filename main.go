package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
        
        portNumber := os.Getenv("WEBSERVER_PORT")

	fmt.Println("Listening on port: " + portNumber + " .....")
	log.Fatal(http.ListenAndServe(":" + portNumber, http.FileServer(http.Dir("./html"))))

}

