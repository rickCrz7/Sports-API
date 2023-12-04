package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ConvertToInt(id string) int {
	intID, _ := strconv.Atoi(id)
	return intID
}
