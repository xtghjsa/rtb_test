package api

import (
	"database/sql"
	"test_project/internal/api/handler"
	"test_project/internal/api/usecase"
	"test_project/internal/repository"

	"github.com/gin-gonic/gin"
)

func StartServer(serverHost, serverPort string, db *sql.DB) error {
	//Initialize repository
	repo := &repository.PostgresConnection{DB: db}
	//Initialize usecases
	dspUsecase := &usecase.DspUsecase{Repo: repo}
	//Initialize handlers
	dsp1Handler := &handler.Dsp1Handler{Usecase: dspUsecase}
	dsp2Handler := &handler.Dsp2Handler{Usecase: dspUsecase}
	dsp3Handler := &handler.Dsp3Handler{Usecase: dspUsecase}

	r := gin.Default()
	//DSP handlers
	r.POST("/dsp1", dsp1Handler.Dsp1)
	r.POST("/dsp2", dsp2Handler.Dsp2)
	r.POST("/dsp3", dsp3Handler.Dsp3)
	//SSP handler
	r.POST("/ssp", handler.Ssp)

	if err := r.Run(serverHost + ":" + serverPort); err != nil {
		return err
	}
	return nil
}
