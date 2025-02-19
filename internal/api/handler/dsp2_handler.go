package handler

import (
	"net/http"
	"test_project/internal/api/request"
	"test_project/internal/api/response"
	"test_project/internal/api/usecase"
	"test_project/internal/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type Dsp2Handler struct {
	Usecase *usecase.DspUsecase
}

// Gets AdCondition, responses with adSpecs to SSP
func (u *Dsp2Handler) Dsp2(c *gin.Context) {
	time.Sleep(500 * time.Millisecond)

	var AdCondition request.DspRequest

	if err := c.ShouldBindJSON(&AdCondition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := u.Usecase.Exec(entities.Ad{
		AdCondition: AdCondition.AdCondition,
		DspID:       "2",
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
		Price:       9999,
	}
	c.JSON(http.StatusOK, DspResponse)
}
