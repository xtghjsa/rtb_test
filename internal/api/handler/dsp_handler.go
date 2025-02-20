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

	var conditions request.DspRequest

	if err := c.ShouldBindJSON(&conditions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, price, err := u.Usecase.Exec(entities.Ad{
		AdCondition: conditions.AdCondition,
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
	time.Sleep(time.Duration(conditions.Delay) * time.Millisecond)
	log.Println("DSP: "+c.Param("id")+" Price: ", price)
	c.JSON(http.StatusOK, DspResponse)
}
