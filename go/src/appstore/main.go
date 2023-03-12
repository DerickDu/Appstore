package main

import (
	"appstore/backend"
	"appstore/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("Hello World")
	fmt.Println("started-service")
	backend.InitElasticsearchBackend()
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
