package usecase

import "test_project/internal/entities"

type DspUsecaseInt interface {
	GetAd(ad entities.Ad) (entities.Ad, error)
}

type DspUsecase struct {
	Repo DspUsecaseInt
}

func (u *DspUsecase) Exec(ad entities.Ad) (entities.Ad, error) {
	return u.Repo.GetAd(ad)
}
