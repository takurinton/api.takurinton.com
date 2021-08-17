package controller

import (
	"net/http"
	"portfolio/pagination"
	"portfolio/service"
	"portfolio/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetDairyReport(c *gin.Context) {
	d := service.Dairy{}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	r, err := d.GetAllDairy(page)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		next, prev := pagination.GetNippoParams(r.Page, r.TotalPage, r.NextPage, r.PrevPage)
		c.JSONP(http.StatusOK, gin.H{
			"next":    next,
			"prev":    prev,
			"results": r.Records,
		})
	}
}

func CreateDailyreport(c *gin.Context) {
	d := service.Dairy{}
	name := c.PostForm("contents")
	_pubDate := c.PostForm("pub_date")
	pubDate := utils.StringToTime(_pubDate)

	daily, err := d.CreateDailyreport(name, pubDate)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"daily": daily})
	}
}

func GetDetailDairyReport(c *gin.Context) {
	d := service.Dairy{}
	id := c.Param("id")
	results, err := d.GetDetailDairy(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		var comment []string
		c.JSONP(http.StatusOK, gin.H{
			"post":    results,
			"comment": comment,
		})
	}
}

func UpdateDairyReport(c *gin.Context) {
	d := service.Dairy{}
	id := c.Param("id")
	contents := c.PostForm("contents")

	if err := d.UpdateDailyreport(id, contents); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"id":       id,
			"contents": contents,
		})
	}
}

// func DeleteDetailDairyReport(c *gin.Context) {}
