package main

import (
	"log"
	"net/http"

	"github.com/Labastidaa/go-blckchn/internal/api"
	"github.com/Labastidaa/go-blckchn/internal/router"

	"github.com/rs/cors"
)

func main() {
	apiClient := api.NewCoinMarketCapClient(nil)
	r := router.SetupRouter(apiClient)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://app-latest-kco7.onrender.com", "https://www.glvz-ds.com"},
		AllowedMethods:   []string{http.MethodGet, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CMC_PRO_API_KEY"},
		AllowCredentials: true,
		Debug:            true,
	})

	// Apply CORS middleware to our router
	handler := c.Handler(r)

	log.Println("Server starting on port :8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
