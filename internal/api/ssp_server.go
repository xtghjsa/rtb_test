package api

import (
	"test_project/internal/api/handler"
	"test_project/internal/api/usecase"
	"test_project/internal/auction"
	"test_project/internal/repository"
	"test_project/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func StartSSP(SSPHost, SSPPort string, kfkWriter *kafka.Writer, cfg utils.AuctionConfig) error {

	r := gin.Default()
	//Tracking handler initialization
	repo := &repository.KafkaWriter{Writer: kfkWriter}
	trackingUsecase := &usecase.TrackingUsecase{Repo: repo}
	trackingHandler := &handler.TrackingHandler{Usecase: trackingUsecase}
	//Auction handler initialization
	auctionConfig := &auction.AuctionService{Cfg: &cfg}
	auctionUsecase := &usecase.AuctionUsecase{Cfg: auctionConfig}
	auctionHandler := &handler.AuctionHandler{Usecase: auctionUsecase}

	//SSP handlers
	r.POST("/ssp", auctionHandler.Auction)
	r.GET("/tracking", trackingHandler.Tracking)

	if err := r.Run(SSPHost + ":" + SSPPort); err != nil {
		return err
	}
	return nil
}
