package utils

import (
	"api/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// DelDuplicate удаляет дубликаты
func DelDuplicate(links models.TaskRequest) map[string]bool {
	var data = make(map[string]bool)

	for _, link := range links.Links {
		data[link] = true
	}
	return data
}

func GetStatus(url string) error {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	_, err := client.Get(url)
	return err
}

func AddStorage(cache *models.Cache) error {

	cache.Mu.Lock()
	defer cache.Mu.Unlock()
	if len(cache.List) == 0 {
		return nil
	}

	for k, v := range cache.List {
		linksResponse := models.TaskResponse{
			Links:    map[string]string{},
			LinksNum: k,
		}
		for _, link := range v {
			url := link

			if !strings.HasPrefix(link, "https://") {
				url = "https://" + link
			}
			err := GetStatus(url)
			if err != nil {
				linksResponse.Links[link] = "not available"
			} else {
				linksResponse.Links[link] = "available"
			}
		}

		dataJson, err := json.Marshal(linksResponse)
		if err != nil {
			return err
		}

		_, err = http.Post("http://localhost:8080/restore", "application/json", bytes.NewBuffer(dataJson))
		if err != nil {
			return err
		}
	}
	return nil
}
