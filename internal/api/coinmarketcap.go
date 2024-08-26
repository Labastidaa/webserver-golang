package api

import (
	"context"
	"encoding/json"
	"fmt"

	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/Labastidaa/go-blckchn/internal/models"
)

type CoinMarketCapClient struct {
	Api CoinMarketCapAPI
}

type CoinMarketCapAPI interface {
	FetchCryptoListings() (*models.APIResponse, error)
}

var (
	once     sync.Once
	instance *CoinMarketCapClient
)

func NewCoinMarketCapClient(api CoinMarketCapAPI) *CoinMarketCapClient {

	once.Do(func() {
		if api == nil {
			api = &defaultCoinMarketCapAPI{
				BaseURL:    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest",
				APIKey:     os.Getenv("COINMARKETCAP_API_KEY"),
				HTTPClient: &http.Client{},
			}
		}
		instance = &CoinMarketCapClient{
			Api: api,
		}
	})
	return instance
}

type defaultCoinMarketCapAPI struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func (api *defaultCoinMarketCapAPI) FetchCryptoListings() (*models.APIResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", api.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "USD")
	q.Add("sort", "market_cap")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", api.APIKey)

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s (%d)", string(body), resp.StatusCode)
	}

	var response models.APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing JSON: %s", err)
		return nil, err
	}

	return &response, nil
}

func (client *CoinMarketCapClient) FetchCryptoListings() (*models.APIResponse, error) {
	return client.Api.FetchCryptoListings()
}

var _ CoinMarketCapAPI = &defaultCoinMarketCapAPI{}
