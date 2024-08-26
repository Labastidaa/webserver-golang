package router

import (
	"github.com/Labastidaa/go-blckchn/internal/api"
	"github.com/Labastidaa/go-blckchn/internal/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(apiClient *api.CoinMarketCapClient) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/crypto", handler.CryptoHandler(apiClient)).Methods("GET")
	return r
}
