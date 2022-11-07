package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Data float64 `json:"data"`
}

// handles POST request of route "/api/calculate-integral"
func HandleCalculateIntegral(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var funcParams FuncParams
	json.NewDecoder(r.Body).Decode(&funcParams)

	result := Result{}
	result.Data = CalculateIntegralWithPrecision(&funcParams)
	json.NewEncoder(w).Encode(&result)
}
