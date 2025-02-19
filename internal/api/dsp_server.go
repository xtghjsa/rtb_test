package api

import (
	"database/sql"
	"test_project/internal/api/handler"
	"test_project/internal/api/usecase"
	"test_project/internal/repository"

	"github.com/gin-gonic/gin"
)

func StartDSP(DSPHost, DSPPort string, db *sql.DB) error {
	//Initialize repository
	repo := &repository.PostgresConnection{DB: db}
	//Initialize usecases
	dspUsecase := &usecase.DspUsecase{Repo: repo}
	//Initialize handlers
	dspHandler := &handler.DspHandler{Usecase: dspUsecase}

	r := gin.Default()
	//DSP handlers
	r.POST("/dsp/:id", dspHandler.Dsp)

	if err := r.Run(DSPHost + ":" + DSPPort); err != nil {
		return err
	}
	return nil
}
