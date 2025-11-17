package router

import (
	"api/internal/api/handlers"
	"net/http"
)

func TaskRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", handlers.GetStatusServer)
	return mux
}
