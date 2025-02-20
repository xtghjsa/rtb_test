package usecase

import (
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/utils"
)

type SspUsecaseInt interface {
	SSPAuctionService(condition request.DspRequest) response.SspResponse
}

func SspExec(conditionsSSP request.SspRequest) response.SspResponse {
	conditionsDSP := request.DspRequest{
		AdCondition: conditionsSSP.AdCondition,
	}
	return utils.SSPAuctionService(conditionsDSP)
}
