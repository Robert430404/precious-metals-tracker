package pricing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func GetPricingRepository() (*PricingRepository, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	apiKey := config.RuntimeFlags.GoldAPIKey

	repository := &PricingRepository{
		ApiKey:     apiKey,
		ApiBaseUrl: "https://www.goldapi.io",
		HttpClient: &http.Client{},
	}

	return repository, nil
}

func (self *PricingRepository) LoadCachedBytes(fileName string) ([]byte, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	configPath := config.ConfigPath
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
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	configPath := config.ConfigPath

	return os.WriteFile(configPath+"/"+fileName, payload, 0644)
}

func (self *PricingRepository) GetSilverSpot() float64 {
	cachedResponseBytes, _ := self.LoadCachedBytes("silver-price-response.json")

	var cachedResponse PriceResponse
	json.Unmarshal(cachedResponseBytes, &cachedResponse)

	oneDayAgo := time.Now().AddDate(0, 0, -1)
	cacheOlderThanOneDay := cachedResponse.Timestamp < oneDayAgo.Unix()
	if !cacheOlderThanOneDay {
		return cachedResponse.Price
	}

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

	body, err5 := io.ReadAll(resp.Body)
	if err5 != nil {
		fmt.Println("there was a problem retrieving the new spot price, using cache")
		return cachedResponse.Price
	}

	var httpResponse PriceResponse
	err6 := json.Unmarshal(body, &httpResponse)
	if err6 != nil {
		return cachedResponse.Price
	}

	self.WriteCacheBytes("silver-price-response.json", body)

	return httpResponse.Price
}

func (self *PricingRepository) GetGoldSpot() float64 {
	cachedResponseBytes, _ := self.LoadCachedBytes("gold-price-response.json")

	var cachedResponse PriceResponse
	json.Unmarshal(cachedResponseBytes, &cachedResponse)

	oneDayAgo := time.Now().AddDate(0, 0, -1)
	cacheOlderThanOneDay := cachedResponse.Timestamp < oneDayAgo.Unix()
	if !cacheOlderThanOneDay {
		return cachedResponse.Price
	}

	req, err3 := http.NewRequest("GET", self.ApiBaseUrl+"/api/XAU/USD", nil)
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

	body, err5 := io.ReadAll(resp.Body)
	if err5 != nil {
		fmt.Println("there was a problem retrieving the new spot price, using cache")
		return cachedResponse.Price
	}

	var httpResponse PriceResponse
	err6 := json.Unmarshal(body, &httpResponse)
	if err6 != nil {
		return cachedResponse.Price
	}

	self.WriteCacheBytes("gold-price-response.json", body)

	return httpResponse.Price
}
