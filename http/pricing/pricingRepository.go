package pricing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/robert430404/precious-metals-tracker/config"
)

type PricingRepository struct {
	ApiKey     string
	ApiBaseUrl string
	HttpClient *http.Client
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

func (self *PricingRepository) WriteCacheBytes(fileName string, payload []byte) error {
	configPath := config.GetConfig().ConfigPath
	return os.WriteFile(configPath+"/"+fileName, payload, 0644)
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

	oneDayAgo := time.Now().AddDate(0, 0, -1)
	cacheOlderThanOneDay := cachedResponse.Timestamp < oneDayAgo.Unix()
	if cacheOlderThanOneDay {
		req, err3 := http.NewRequest("GET", self.ApiBaseUrl+"/api/XAG/USD", nil)
		if err3 != nil {
			fmt.Println("there was a problem retrieving the new spot price, using cache")
			return cachedResponse.Price
		}

		req.Header.Add("x-access-token", self.ApiKey)

		resp, err4 := self.HttpClient.Do(req)
		if err4 != nil {
			fmt.Println("there was a problem retrieving the new spot price, using cache")
			return cachedResponse.Price
		}

		defer resp.Body.Close()

		body, err5 := ioutil.ReadAll(resp.Body)
		if err5 != nil {
			fmt.Println("there was a problem retrieving the new spot price, using cache")
			return cachedResponse.Price
		}

		var httpResponse PriceResponse
		err6 := json.Unmarshal(body, &httpResponse)
		if err6 != nil {
			return cachedResponse.Price
		}

		self.WriteCacheBytes("price-response.json", body)

		return httpResponse.Price
	}

	return cachedResponse.Price
}

func GetPricingRepository() *PricingRepository {
	apiKey := config.GetConfig().RuntimeFlags.GoldAPIKey

	repository := &PricingRepository{
		ApiKey:     apiKey,
		ApiBaseUrl: "https://www.goldapi.io",
		HttpClient: &http.Client{},
	}

	return repository
}
