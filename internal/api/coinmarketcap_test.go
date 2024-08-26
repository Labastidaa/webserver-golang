package api

import (
	"testing"

	"github.com/Labastidaa/go-blckchn/internal/api/mocks"
	"github.com/Labastidaa/go-blckchn/internal/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFetchCryptoListings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	falseVal := false
	expectedResponse := &models.APIResponse{
		Data: []models.CryptoData{
			{
				ID:                1,
				Name:              "Bitcoin",
				Symbol:            "BTC",
				Slug:              "bitcoin",
				CMCRank:           5,
				NumMarketPairs:    500,
				CirculatingSupply: 16950100,
				TotalSupply:       16950100,
				MaxSupply:         21000000,
				InfiniteSupply:    &falseVal,
				LastUpdated:       "2018-06-02T22:51:28.209Z",
				DateAdded:         "2013-04-28T00:00:00.000Z",
				Tags:              []string{"mineable"},
				Platform:          nil,
				Quote: map[string]models.QuoteDetails{
					"USD": {
						Price:                 9283.92,
						Volume24H:             7155680000,
						PercentChange1H:       -0.152774,
						PercentChange24H:      0.518894,
						PercentChange7D:       0.986573,
						MarketCap:             852164659250.2758,
						FullyDilutedMarketCap: 952835089431.14,
						LastUpdated:           "2018-08-09T22:53:32.000Z",
					},
					"BTC": {
						Price:                 1,
						Volume24H:             772012,
						MarketCap:             17024600,
						FullyDilutedMarketCap: 952835089431.14,
						LastUpdated:           "2018-08-09T22:53:32.000Z",
					},
				},
			},
		},
		Status: models.StatusInfo{
			Timestamp:    "2018-06-02T22:51:28.209Z",
			ErrorCode:    0,
			ErrorMessage: "",
			Elapsed:      10,
			CreditCount:  1,
		},
	}

	mockApi := mocks.NewMockCoinMarketCapAPI(ctrl)
	mockApi.EXPECT().FetchCryptoListings().Return(expectedResponse, nil)

	client := NewCoinMarketCapClient(mockApi)

	response, err := client.FetchCryptoListings()
	if err != nil {
		t.Fatal("FetchCryptoListings failed:", err)
	}

	assert.NoError(t, err, "FetchCryptoListings should not return an error")
	assert.Equal(t, expectedResponse, response, "Response should match the expected API response")

}
