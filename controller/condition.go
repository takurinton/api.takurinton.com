package controller

import (
	"net/http"
	"portfolio/service"

	"github.com/gin-gonic/gin"
)

func GetAllCondition(c *gin.Context) {
	m := service.MaririntonCondition{}
	results, err := m.GetCondition("")
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusOK, results)
	}
}
