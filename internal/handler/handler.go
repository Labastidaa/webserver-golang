package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Labastidaa/go-blckchn/internal/api"
)

func CryptoHandler(client *api.CoinMarketCapClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiResponse, err := client.FetchCryptoListings()
		if err != nil {
			log.Printf("Error fetching crypto listings: %v", err)
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(apiResponse)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}

}
