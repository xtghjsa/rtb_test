package auction

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/utils"
	"time"
)

func SSPAuctionService(condition request.DspRequest, endpoints []string) response.SspResponse {

	resultsChannel := make(chan response.DspResponse, len(endpoints))
	for _, url := range endpoints {
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

	results := make([]response.DspResponse, len(endpoints))
	timeLimit := time.After(300 * time.Millisecond)
loop:
	for i := 0; i < len(endpoints); i++ {
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
	response := response.SspResponse{
		DspID:  mostExpensive.DspID,
		AdName: mostExpensive.AdName,
	}
	return response
}
