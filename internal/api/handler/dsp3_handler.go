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

type Dsp3Handler struct {
	Usecase *usecase.DspUsecase
}

// Gets AdCondition, responses with adSpecs to SSP
func (u *Dsp3Handler) Dsp3(c *gin.Context) {
	time.Sleep(50 * time.Millisecond)

	var AdCondition request.DspRequest

	if err := c.ShouldBindJSON(&AdCondition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := u.Usecase.Exec(entities.Ad{
		AdCondition: AdCondition.AdCondition,
		DspID:       "3",
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
		Price:       500,
	}
	c.JSON(http.StatusOK, DspResponse)
}
