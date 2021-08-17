package router

import (
	"net/http"
	"portfolio/controller"
	"portfolio/gql"
	"portfolio/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           time.Duration
}

func Router() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	r.Use(cors.New(config))

	r.GET("/", controller.Root) // fake
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(http.StatusOK, "アプリケーションの鯖生きてるぽよぽよ〜！！")
	})

	dairyrouter := r.Group("/dairyreport")
	{
		router := dairyrouter.Group("/v1")
		{
			router.GET("/", controller.GetDairyReport)
			router.GET("/:id", controller.GetDetailDairyReport)
		}
	}

	blogrouter := r.Group("/blog")
	{
		router := blogrouter.Group("/v1")
		{
			router.GET("/", controller.GetPosts)
			router.GET("/rss", controller.GetAllPosts)
			router.GET("/post/:id", controller.GetPost)
			router.POST("/comment/:id", controller.PostComment)
		}
	}

	portfoliorouter := r.Group("/portfolio")
	{
		router := portfoliorouter.Group("v1") // 多分一生v1 (やる気ないので)
		{
			router.GET("/", controller.GetPortfolio)
			router.GET("mine101content", controller.Get101Content)
			router.POST("/contact", controller.PostContact)
		}
	}

	userrouter := r.Group("/admin/user")
	{
		userrouter.POST("/register", controller.CreateUser)
		userrouter.POST("/login", controller.Login)
	}

	marinarouter := r.Group("/marina")
	{
		marinarouter.GET("/comment", controller.GetOmedetou)
		marinarouter.GET("/congrats", controller.GetCongrats)
		marinarouter.POST("/comment", controller.PostOmedetou)
		marinarouter.POST("/congrats", controller.PostCongrats)
	}

	adminrouter := r.Group("/admin")
	adminrouter.Use(middleware.AuthMiddlewere)
	{
		router := adminrouter
		{
			router.GET("/users", controller.GetUsers)
			router.POST("/logout", controller.Logout)

			router.GET("/blog", controller.GetAllPostsReverse)
			router.POST("/blog", controller.CreatePost)
			router.GET("/blog/post/:id", controller.GetPostAdmin)
			router.PATCH("/blog/post/:id", controller.UpdateDetailPost)
			router.DELETE("/blog/post/:id", controller.DeleteDetailPost)

			router.GET("/blog/category", controller.GetCategory)
			router.POST("/blog/category", controller.CreateCategory)
			router.GET("/blog/category/:id", controller.GetDetailCategory)
			router.PATCH("/blog/category/:id", controller.UpdateCategory)
			router.DELETE("/blog/category/:id", controller.DeleteCategory)

			router.GET("/daily", controller.GetDairyReport)
			router.POST("/daily", controller.CreateDailyreport)
			router.GET("/daily/post/:id", controller.GetDetailDairyReport)
			router.PATCH("/daily/post/:id", controller.UpdateDairyReport)
		}
	}

	slackrouter := r.Group("/slack")
	{
		slackrouter.GET("/", controller.Hoge)
		slackrouter.POST("/", controller.HelloWorld)
		slackrouter.POST("/condition", controller.Condition)
	}

	conditionrouter := r.Group("/condition")
	conditionrouter.Use(middleware.AuthMiddlewere)
	{
		conditionrouter.GET("", controller.GetAllCondition)
	}

	r.POST("/graphql", gql.Response)
	r.POST("/graphql/condition", gql.ConditionResponse)

	http.ListenAndServe(":8080", r)

	return r
}
