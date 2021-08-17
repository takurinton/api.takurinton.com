package controller

import (
	"net/http"
	"portfolio/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Get101Content(c *gin.Context) {
	h := service.Get101Content{}

	results, err := h.Get101Content()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, results)
	}
}
