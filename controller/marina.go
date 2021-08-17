package controller

import (
	"net/http"
	"portfolio/model"
	"portfolio/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetOmedetou(c *gin.Context) {
	h := service.Marina{}

	results, err := h.GetOmedetou()

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

func GetCongrats(c *gin.Context) {
	h := service.Marina{}

	results, err := h.GetCongrats()
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

func PostOmedetou(c *gin.Context) {
	h := service.Marina{}

	var omedetou model.MarinaMarinaOmedetou
	c.BindJSON(&omedetou)

	err := h.PostOmedetou(omedetou.Name, omedetou.Sentence)
	if err != nil {
	}

}

func PostCongrats(c *gin.Context) {
	h := service.Marina{}

	var omedetouCount model.MarinaMarinaOmedetouCount
	c.BindJSON(&omedetouCount)

	err := h.PostCongrats(omedetouCount.Count)
	if err != nil {
	}
}
