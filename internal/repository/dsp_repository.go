package repository

import (
	"database/sql"
	"test_project/internal/entities"
)

type PostgresConnection struct {
	DB *sql.DB
}

// Only for testing purposes
func AddTestAds(DB *sql.DB) error {
	dspIds := []string{"1", "2", "3"}
	for _, dspId := range dspIds {
		_, err := DB.Exec("INSERT INTO ads (dsp_id, ad_name, ad_condition) VALUES ($1, $2, $3)", dspId, "name"+dspId, "condition")
		if err != nil {
			return err
		}
	}
	return nil
}

// Get ad from database depending on ad_condition and dsp_id
func (p *PostgresConnection) GetAd(ad entities.Ad) (entities.Ad, error) {
	err := p.DB.QueryRow("SELECT id, ad_name FROM ads WHERE ad_condition = $1 AND dsp_id = $2", ad.AdCondition, ad.DspID).Scan(&ad.ID, &ad.AdName)
	if err != nil {
		return ad, err
	}
	return ad, nil
}
