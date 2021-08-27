package controller

import (
	"net/http"
	"portfolio/model"
	"portfolio/service"

	"github.com/gin-gonic/gin"
)

func Analytics(c *gin.Context) {
	h := service.Analytics{}
	var analytics model.Analytics
	if err := c.BindJSON(&analytics); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	h.AddAnalytics(analytics)
}
