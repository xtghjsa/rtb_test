package usecase

import (
	"test_project/internal/api/request"
	"test_project/internal/api/response"
)

type AuctionUsecaseInt interface {
	SSPAuctionService(condition request.DspRequest) response.SspResponse
}

type AuctionUsecase struct {
	Cfg AuctionUsecaseInt
}

func (c *AuctionUsecase) AuctionExec(conditionsSSP request.SspRequest) response.SspResponse {
	conditionsDSP := request.DspRequest{
		AdCondition: conditionsSSP.AdCondition,
	}
	return c.Cfg.SSPAuctionService(conditionsDSP)
}
