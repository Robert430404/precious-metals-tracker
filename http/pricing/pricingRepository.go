package pricing

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/robert430404/precious-metals-tracker/config"
)

type PricingRepository struct {
	ApiKey string
	ApiBaseUrl string
}

type PriceResponse struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"open_price"`
}

func (self *PricingRepository) LoadCachedBytes(fileName string) ([]byte, error) {
	configPath := config.GetConfig().ConfigPath
	rawCachedBytes, err := os.Open(configPath + "/" + fileName)
	if err != nil {
		return nil, errors.New("could not load the cached response")
	}

	defer rawCachedBytes.Close()

	cachedBytes, err1 := io.ReadAll(rawCachedBytes)
	if err1 != nil {
		return nil, errors.New("could not read the cached response")
	}

	return cachedBytes, nil
}

func (self *PricingRepository) WriteCacheBytes(fileName string, payload string) error {
	return nil	
}

func (self *PricingRepository) GetSilverSpot() float64 {
	cachedResponseBytes, err := self.LoadCachedBytes("price-response.json")
	if err != nil {
		return 0
	}

	var cachedResponse PriceResponse
	err2 := json.Unmarshal(cachedResponseBytes, &cachedResponse)
	if err2 != nil {
		return 0
	}

	return cachedResponse.Price
}

func GetPricingRepository() *PricingRepository {
	apiKey := config.GetConfig().RuntimeFlags.GoldAPIKey

	repository := &PricingRepository{
		ApiKey: apiKey,
		ApiBaseUrl: "https://www.goldapi.io",
	}

	return repository
}