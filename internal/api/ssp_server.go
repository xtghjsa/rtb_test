package api

import (
	"test_project/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func StartSSP(SSPHost, SSPPort string) error {

	r := gin.Default()
	//SSP handler
	r.POST("/ssp", handler.Ssp)

	if err := r.Run(SSPHost + ":" + SSPPort); err != nil {
		return err
	}
	return nil
}
