package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/storage"

	"github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, found := storage.GetPoints(id)

	if !found {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
