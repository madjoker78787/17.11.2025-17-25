package utils

import (
	"api/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func SendToTask(links models.DataResponse) (*http.Response, bool) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	data, err := json.Marshal(links)
	resp, err := client.Post("http://localhost:8081/check", "application/json", bytes.NewReader(data))
	if err != nil {
		return &http.Response{}, false
	}
	return resp, true
}
