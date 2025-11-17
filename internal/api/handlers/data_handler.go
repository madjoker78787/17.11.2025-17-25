package handlers

import (
	"api/internal/models"
	"api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func SetData(w http.ResponseWriter, r *http.Request, memory *models.Storage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var links = models.DataResponse{}
	err := json.NewDecoder(r.Body).Decode(&links)
	if err != nil {
		http.Error(w, "Invalid data request", http.StatusBadRequest)
		return
	}

	memory.SetNextID()

	resp, ok := utils.SendToTask(links)
	if !ok {
		fmt.Println("Записали")
		memory.Cached.Set(memory.NextID, links.Links)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("Data has been received, you can get the result by index: %d", memory.NextID)))
		fmt.Println("memory.Cached.List", memory.Cached.List)
		return
	}
	defer resp.Body.Close()

	var dataRequest models.DataRequest
	err = json.NewDecoder(resp.Body).Decode(&dataRequest)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRequest.LinksNum = memory.NextID

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dataRequest)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	item := models.Item{
		Links: dataRequest.Links,
	}
	memory.Set(item)
	fmt.Println("memory.StorageList", memory.StorageList)
	fmt.Println("memory.Cached.List", memory.Cached.List)

}

func RestoreData(w http.ResponseWriter, r *http.Request, memory *models.Storage) {
	switch r.Method {
	case http.MethodGet:
		cachedList := memory.Cached.Get()
		err := json.NewEncoder(w).Encode(cachedList)
		if err != nil {
			return
		}
	case http.MethodPost:
		data := models.DataRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			fmt.Println("POST ", err)
			return
		}
		fmt.Println("data.LinksNum", data.LinksNum)
		fmt.Println("data.Links", data.Links)
		item := models.Item{
			Links: data.Links,
		}
		memory.Set(item)
		fmt.Println("memory.StorageList", memory.StorageList)
		fmt.Println("memory.Cached.List", memory.Cached.List)
		delete(memory.Cached.List, data.LinksNum)
		fmt.Println("memory.Cached.List", memory.Cached.List)
	}

}
