package api

import (
	"test_project/internal/api/handler"
	"test_project/internal/api/usecase"
	"test_project/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func StartSSP(SSPHost, SSPPort string, kfkWriter *kafka.Writer) error {

	r := gin.Default()

	repo := &repository.KafkaWriter{Writer: kfkWriter}
	trackingUsecase := &usecase.TrackingUsecase{Repo: repo}
	trackingHandler := &handler.TrackingHandler{Usecase: trackingUsecase}

	//SSP handlers
	r.POST("/ssp", handler.Ssp)
	r.GET("/tracking", trackingHandler.Tracking)

	if err := r.Run(SSPHost + ":" + SSPPort); err != nil {
		return err
	}
	return nil
}
