package handler

import (
	"log"
	"net/http"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/api/usecase"
	"test_project/internal/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type DspHandler struct {
	Usecase *usecase.DspUsecase
}

// Gets AdCondition, responses with adSpecs to SSP
func (u *DspHandler) Dsp(c *gin.Context) {
	time.Sleep(5 * time.Millisecond)

	var AdCondition request.DspRequest

	if err := c.ShouldBindJSON(&AdCondition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, price, err := u.Usecase.Exec(entities.Ad{
		AdCondition: AdCondition.AdCondition,
		DspID:       c.Param("id"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var DspResponse = response.DspResponse{
		ID:          result.ID,
		DspID:       result.DspID,
		AdName:      result.AdName,
		AdCondition: result.AdCondition,
		Price:       price,
	}
	log.Println("DSP: "+c.Param("id")+" Price: ", price)
	c.JSON(http.StatusOK, DspResponse)
}
