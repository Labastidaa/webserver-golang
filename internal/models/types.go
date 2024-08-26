package models

type APIResponse struct {
	Status StatusInfo   `json:"status"`
	Data   []CryptoData `json:"data"`
}

type StatusInfo struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}

type CryptoData struct {
	ID                            int                     `json:"id"`
	Name                          string                  `json:"name"`
	Symbol                        string                  `json:"symbol"`
	Slug                          string                  `json:"slug"`
	CMCRank                       int                     `json:"cmc_rank"`
	NumMarketPairs                int                     `json:"num_market_pairs"`
	CirculatingSupply             float64                 `json:"circulating_supply"`
	TotalSupply                   float64                 `json:"total_supply"`
	MaxSupply                     float64                 `json:"max_supply"`
	InfiniteSupply                *bool                   `json:"infinite_supply"`
	LastUpdated                   string                  `json:"last_updated"`
	DateAdded                     string                  `json:"date_added"`
	Tags                          []string                `json:"tags"`
	Platform                      *interface{}            `json:"platform"`
	SelfReportedCirculatingSupply *float64                `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64                `json:"self_reported_market_cap"`
	Quote                         map[string]QuoteDetails `json:"quote"`
}

type QuoteDetails struct {
	Price                 float64 `json:"price"`
	Volume24H             float64 `json:"volume_24h"`
	VolumeChange24H       float64 `json:"volume_change_24h"`
	PercentChange1H       float64 `json:"percent_change_1h"`
	PercentChange24H      float64 `json:"percent_change_24h"`
	PercentChange7D       float64 `json:"percent_change_7d"`
	MarketCap             float64 `json:"market_cap"`
	MarketCapDominance    float64 `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
	LastUpdated           string  `json:"last_updated"`
}
