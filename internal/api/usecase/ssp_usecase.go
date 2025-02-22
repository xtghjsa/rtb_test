package usecase

import (
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/auction"
)

type SspUsecaseInt interface {
	SSPAuctionService(condition request.DspRequest, endpoints []string) response.SspResponse
}

func AuctionExec(conditionsSSP request.SspRequest) response.SspResponse {
	conditionsDSP := request.DspRequest{
		AdCondition: conditionsSSP.AdCondition,
	}
	dspEndpoints := []string{
		"http://dsp:7070/dsp/1",
		"http://dsp:7070/dsp/2", //change localhost to dsp for docker
		"http://dsp:7070/dsp/3",
	}
	return auction.SSPAuctionService(conditionsDSP, dspEndpoints)
}
