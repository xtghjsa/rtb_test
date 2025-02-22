package handler

import (
	"net/http"
	"test_project/internal/api/request"
	"test_project/internal/api/usecase"

	"github.com/gin-gonic/gin"
)

type AuctionHandler struct {
	Usecase *usecase.AuctionUsecase
}

// Gets AdCondition, sends it to the DSP, gets the response from the DSP, sends it to the user
func (u *AuctionHandler) Auction(c *gin.Context) {
	var condition request.SspRequest

	if err := c.ShouldBindJSON(&condition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := u.Usecase.AuctionExec(condition)

	c.JSON(http.StatusOK, resp)
}
