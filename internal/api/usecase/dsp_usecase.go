package usecase

import (
	"test_project/internal/entities"
	"test_project/internal/utils"
)

type DspUsecaseInt interface {
	GetAd(ad entities.Ad) (entities.Ad, error)
}

type DspUsecase struct {
	Repo DspUsecaseInt
}

func (u *DspUsecase) Exec(ad entities.Ad) (adResp entities.Ad, price int64, bidID string, err error) {
	adResp, err = u.Repo.GetAd(ad)
	if err != nil {
		return adResp, 0, "", err
	}
	price = utils.RandomPrice()
	bidID = utils.GenerateBidID()
	return adResp, price, bidID, nil
}
