package main

import (
	"api/internal/api/router"
	"api/internal/models"
	"api/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://localhost:8080/restore")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	cache := models.Cache{}

	err = json.NewDecoder(resp.Body).Decode(&cache.List)
	if err != nil || len(cache.List) == 0 {
		return
	}

	err = utils.AddStorage(&cache)
	if err != nil {
		return
	}
	cache.Mu.Lock()
	cache.List = make(map[int][]string)
	cache.Mu.Unlock()
}

func main() {

	port := ":8081"

	mux := router.TaskRouter()

	fmt.Println("Task server is running on port: ", port)

	log.Fatal(http.ListenAndServe(port, mux))
}
