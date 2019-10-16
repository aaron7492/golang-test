package main

import (
	"log"
	"net/http"
)

func main() {
	//levantar web server
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
