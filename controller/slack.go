package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

func Hoge(c *gin.Context) {
	res := gin.H{
		"status": "ok",
	}
	c.JSONP(http.StatusCreated, res)
}

func setToken() *slack.Client {
	token := "xoxb-1978524669680-2332177912432-oniubBCI7jm3Fvfe8uligKmG"
	client := slack.New(token)
	return client
}

// テスト用のエンドポイント
func HelloWorld(c *gin.Context) {
	client := setToken()

	_, _, err := client.PostMessage("condition", slack.MsgOptionText("<@U01UFA1TP8R>", false))
	if err != nil {
		panic(err)
	}
}

func Condition(c *gin.Context) {
	var test interface{}
	if err := c.BindJSON(&test); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSONP(http.StatusOK, test)
	return
}
