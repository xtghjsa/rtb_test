package handler

import (
	"net/http"
	"test_project/internal/api/usecase"

	"github.com/gin-gonic/gin"
)

type TrackingHandler struct {
	Usecase *usecase.TrackingUsecase
}

func (u *TrackingHandler) Tracking(c *gin.Context) {
	data := c.Query("value")
	eventType := c.Query("event")
	ok := u.Usecase.TrackingExec(data, eventType)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't decode data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
