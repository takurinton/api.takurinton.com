package controller

import (
	"net/http"
	"portfolio/service"

	"portfolio/model"

	"github.com/gin-gonic/gin"
)

func GetPortfolio(c *gin.Context) {
	h := service.Portfolio{}
	results, err := h.GetPortfolio()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"intern": results.Intern,
			"skill":  results.Skill,
			"made":   results.Made,
			"mine":   results.Mine,
		})
	}
}

func PostContact(c *gin.Context) {
	h := service.Portfolio{}
	var contact model.BlogappContact
	c.BindJSON(&contact)
	if err := h.PostContact(contact.Name, contact.Mail, contact.Contents); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, contact)
	}
}
