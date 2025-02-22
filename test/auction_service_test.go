package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/auction"
	"testing"
	"time"
)

func testServer(dspID, adName string, price int64, delay time.Duration) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		resp := response.DspResponse{
			DspID:  dspID,
			AdName: adName,
			Price:  price,
		}
		json.NewEncoder(w).Encode(resp)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestSSPAuctionService(t *testing.T) {
	//Fastest response but low price
	ts1 := testServer("dsp1", "test_ad", 40, 11*time.Millisecond)
	defer ts1.Close()
	//Exceeds the time limit
	ts2 := testServer("dsp2", "test_ad", 555, 444*time.Millisecond)
	defer ts2.Close()
	//Slow response that matches the time limit but has the highest price
	ts3 := testServer("dsp3", "test_ad", 250, 111*time.Millisecond)
	defer ts3.Close()

	endpoints := []string{ts1.URL, ts2.URL, ts3.URL}

	testRequest := request.DspRequest{
		AdCondition: "condition",
	}

	result := auction.SSPAuctionService(testRequest, endpoints)

	if result.DspID != "dsp3" {
		t.Errorf("expected dsp3 but got %s", result.DspID)
	}
}
