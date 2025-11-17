package main

import (
	"api/internal/api/router"
	"api/internal/models"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cache := models.NewCache()
	memory := models.NewStorage(cache)
	port := ":8080"

	mux := router.Router(memory)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Println("Storage server is running on port: ", port)

	log.Fatal(server.ListenAndServe())
}
