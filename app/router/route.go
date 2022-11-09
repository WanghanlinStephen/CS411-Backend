package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"pro/app/http/v1/server"
	"pro/app/middle"
	"pro/app/socket"
	"pro/config"
)

func router(route *gin.Engine) *gin.Engine {

	//socket服务器
	route.GET("/ws", socket.Run)
	//route.GET("/ws/ping", socket.Ping)

	v1 := route.Group("/v1")
	//遊客操作，无需登录
	visitorAPI := v1.Group("/api")
	{
		visitorAPI.GET("test", server.Test)
		visitorAPI.GET("advanceQuery", server.Advance)
		//delete
		visitorAPI.POST("delete", server.Delete)
		//update
		visitorAPI.POST("update", server.Update)
	}

	return route
}

func RouteInit() *gin.Engine {
	if config.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	route := gin.New()
	if config.Mode == "dev" {
		route.Use(gin.Logger())
	}
	route.Use(gin.Recovery()) // 捕捉异常
	route.Use(middle.Access)
	//route.Use(Cors())
	//https config

	//route.Use(TlsHandler())

	return router(route)
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Max-Age", "36000")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, content-type, Accept,Authorization,authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Header("Content-Type", "application/json;charset=utf-8")
			//c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}
		c.Next()
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}
