package handlers

import (
	"api/internal/models"
	"api/pkg/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func GetStatusServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not available", http.StatusBadGateway)
		return
	}
	var links models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&links); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	linksResponse := models.TaskResponse{
		Links: map[string]string{},
	}
	for _, link := range links.Links {
		url := link

		if !strings.HasPrefix(link, "https://") {
			url = "https://" + link
		}

		err := utils.GetStatus(url)
		if err != nil {
			linksResponse.Links[link] = "not available"
		} else {
			linksResponse.Links[link] = "available"
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(linksResponse)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var listNum []int
	if err := json.NewDecoder(r.Body).Decode(&listNum); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

}
