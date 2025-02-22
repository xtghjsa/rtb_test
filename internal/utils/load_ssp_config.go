package utils

import (
	"encoding/json"
	"io"
	"os"
)

type AuctionConfig struct {
	DSPUrls         []string `json:"dsp_urls"`
	MaxResponseTime string   `json:"max_response_time"`
	TrackingURL     string   `json:"tracking_url"`
}

func LoadSSPConfig() (AuctionConfig, error) {
	file, err := os.Open("/app/config/ssp_config.json")
	if err != nil {
		return AuctionConfig{}, err
	}
	defer file.Close()

	cfgBytes, err := io.ReadAll(file)
	if err != nil {
		return AuctionConfig{}, err
	}

	var cfg AuctionConfig
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		return AuctionConfig{}, err
	}
	return cfg, nil

}
