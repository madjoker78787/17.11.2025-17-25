package router

import (
	"api/internal/api/handlers"
	"api/internal/models"
	"net/http"
)

func Router(data *models.Storage) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		handlers.SetData(writer, request, data)
	})

	mux.HandleFunc("/restore", func(writer http.ResponseWriter, request *http.Request) {
		handlers.RestoreData(writer, request, data)
	})

	return mux
}
