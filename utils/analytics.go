package utils

import (
	"portfolio/model"

	"github.com/gin-gonic/gin"
)

func AddAccess(c *gin.Context) (Anal model.Analytics) {
	Anal.Domain = c.Request.Host
	Anal.UA = c.Request.UserAgent()
	Anal.Path = c.Request.URL.Path
	return
}
