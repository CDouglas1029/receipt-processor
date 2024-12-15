package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/storage"
	"receipt-processor-challenge/utils"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	println(json.NewDecoder(r.Body))
	err := json.NewDecoder(r.Body).Decode(&receipt)
	println(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt format. Please verify input", http.StatusBadRequest)
		return
	}

	id := utils.GenerateID()
	points := models.CalculatePoints(receipt)

	storage.StoreReceipt(id, points)

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
