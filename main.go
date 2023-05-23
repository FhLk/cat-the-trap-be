package main

import (
	"cat-the-trap-back-end/controller"
	"cat-the-trap-back-end/midldleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/api/authen", controller.Authen)

	protected := router.Group("/", midldleware.Authentication)
	protected.POST("/api/setup", controller.Setup)
	protected.POST("/api/play", controller.Play)
	protected.POST("/api/reset", controller.Reset)
	protected.POST("/api/time", controller.Time)

	router.Run("192.168.1.232:8080")
	//router.Run("0.0.0.0:3051")
}
