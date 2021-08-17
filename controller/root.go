package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.JSONP(http.StatusOK, `ぽよぽよ〜！にゃんっ！`)
}
