
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Fraud int     `json:"fraud"`
	Score float64 `json:"score"`
}

func fraudHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Fraud: 1, Score: 0.91}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/fraud-check", fraudHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Backend running on :8080")
	http.ListenAndServe(":8080", nil)
}
