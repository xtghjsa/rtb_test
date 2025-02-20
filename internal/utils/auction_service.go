package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"time"
)

func SSPAuctionService(condition request.DspRequest) response.SspResponse {
	dspEndpoints := []string{
		"http://dsp:7070/dsp/1",
		"http://dsp:7070/dsp/2",
		"http://dsp:7070/dsp/3",
	}

	results := make([]response.DspResponse, len(dspEndpoints))

	var wg sync.WaitGroup
	wg.Add(len(dspEndpoints))

	for k, url := range dspEndpoints {
		go func(idx int, endpoint string) {
			defer wg.Done()
			condition.Delay = RandomDelay()
			conditionJSON, _ := json.Marshal(condition)
			resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(conditionJSON))
			if err != nil {
				results[idx] = response.DspResponse{}
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				results[idx] = response.DspResponse{}
				return
			}
			var dspResp response.DspResponse
			err = json.Unmarshal(body, &dspResp)
			if err != nil {
				results[idx] = response.DspResponse{}
				return
			}
			results[idx] = dspResp
		}(k, url)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
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
