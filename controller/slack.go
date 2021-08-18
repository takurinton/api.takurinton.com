package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"os"

	"portfolio/model"
	"portfolio/service"

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
	token := os.Getenv("SLACK_TOKEN")
	client := slack.New(token)
	return client
}

// テスト用のエンドポイント
func HelloWorld(c *gin.Context) {
	var test interface{}
         if err := c.BindJSON(&test); err != nil {
                c.Status(http.StatusBadRequest)
                return
        }
        c.JSONP(http.StatusOK, test)
        return
}

func Condition(c *gin.Context) {
	// var test interface{}
	// if err := c.BindJSON(&test); err != nil {
	// 	 c.Status(http.StatusBadRequest)
	// 	 return
	// }
	// c.JSONP(http.StatusOK, test)
	// return 

	var json model.RequestConditionEventsType

	if err := c.BindJSON(&json); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if json.Event.Channel != "C023DBKUGHK" {
		c.Status(http.StatusInternalServerError)
		return
	}

	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	time.LoadLocation("Asia/Tokyo")
	datetime := time.Unix(int64(json.EventTime), 0)
	user := json.Event.User
	text := json.Event.Text

	checkMental := regexp.MustCompile(`精神\d+(\.\d+)?`).Match([]byte(text))
	checkPhysical := regexp.MustCompile(`体調\d+(\.\d+)?`).Match([]byte(text))
	if !checkMental || !checkPhysical {
		c.Status(http.StatusInternalServerError)
		return
	}
	// (精神or体調)+(整数or小数) にマッチする正規表現
	// 精神\d+(\.\d+)?
	// 体調\d+(\.\d+)?
	m := regexp.MustCompile(`精神\d+(\.\d+)?`)
	p := regexp.MustCompile(`体調\d+(\.\d+)?`)
	men := p.ReplaceAllString(text, "")
	phy := m.ReplaceAllString(text, "")
	mr := regexp.MustCompile(`精神`)
	pr := regexp.MustCompile(`体調`)
	_mental := mr.ReplaceAllString(men, "")
	_physical := pr.ReplaceAllString(phy, "")
	mental, _ := strconv.ParseFloat(strings.Replace(_mental, "\n", "", -1), 64)
	physical, _ := strconv.ParseFloat(strings.Replace(_physical, "\n", "", -1), 64)
	_, err := service.CreateCondition(user, mental, physical, datetime)
	if err != nil {
		fmt.Printf("error %v", err)
	}
}
