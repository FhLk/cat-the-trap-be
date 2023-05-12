package main

import (
	"cat-the-trap-back-end/controller"
	"cat-the-trap-back-end/midldleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func divide(x, y float64) float64 {
	if y == 0 {
		// Trigger a runtime error if y is zero
		panic("cannot divide by zero")
	}
	return x / y
}

func main() {
	//result := divide(10, 0)
	//fmt.Println(result)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(config))

	router.POST("/api/authen", controller.Authen)

	protected := router.Group("/", midldleware.Authentication)
	protected.POST("/api/setup", controller.Setup)
	protected.POST("/api/play", controller.Play)
	protected.POST("/api/reset", controller.Reset)
	//protected.POST("/api/win", controller.Win)
	//protected.POST("/api/lose", controller.Lose)

	router.Run("192.168.1.58:8080")
}
