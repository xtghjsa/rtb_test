package auction

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/encoding"
	"test_project/internal/utils"
	"time"
)

type AuctionService struct {
	Cfg *utils.AuctionConfig
}

func (c *AuctionService) SSPAuctionService(condition request.DspRequest) response.SspResponse {

	resultsChannel := make(chan response.DspResponse, len(c.Cfg.DSPUrls))
	for _, url := range c.Cfg.DSPUrls {
		localCondition := condition
		go func(endpoint string, cond request.DspRequest) {
			cond.Delay = utils.RandomDelay()
			conditionJSON, _ := json.Marshal(cond)
			resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(conditionJSON))
			if err != nil {
				resultsChannel <- response.DspResponse{}
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				resultsChannel <- response.DspResponse{}
				return
			}
			var dspResp response.DspResponse
			err = json.Unmarshal(body, &dspResp)
			if err != nil {
				resultsChannel <- response.DspResponse{}
				return
			}
			resultsChannel <- dspResp
		}(url, localCondition)
	}
	maxResponseTimeInt, err := strconv.Atoi(c.Cfg.MaxResponseTime)
	if err != nil {
		return response.SspResponse{}
	}
	results := make([]response.DspResponse, len(c.Cfg.DSPUrls))
	timeLimit := time.After(time.Duration(maxResponseTimeInt) * time.Millisecond)

loop:
	for i := 0; i < len(c.Cfg.DSPUrls); i++ {
		select {
		case result := <-resultsChannel:
			results = append(results, result)
		case <-timeLimit:
			break loop
		}
	}

	mostExpensive := results[0]
	for i := 1; i < len(results); i++ {
		if mostExpensive.Price < results[i].Price {
			mostExpensive = results[i]
		}
	}
	encoded, _ := encoding.Encode(&mostExpensive)
	response := response.SspResponse{
		DspID:                 mostExpensive.DspID,
		Price:                 mostExpensive.Price,
		AdName:                mostExpensive.AdName,
		TrackingClickURL:      c.Cfg.TrackingURL + "?event=click" + "?value=" + encoded,
		TrackingImpressionURL: c.Cfg.TrackingURL + "?event=impression" + "?value=" + encoded,
	}
	return response
}
