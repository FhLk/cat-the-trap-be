package main

import (
	"cat-the-trap-back-end/Algorithm"
	"cat-the-trap-back-end/service"
	"fmt"
)

func main() {

	board, _, _, _ := service.GameSetup()
	start := board[5][5]
	end := board[0][0]
	path := Algorithm.AStar(start, end, board)
	fmt.Println(path)
	//router := gin.Default()
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:5173"}
	//config.AllowHeaders = []string{"Origin, X-Requested-With, Content-Type, Accept"}
	//config.AllowMethods = []string{"GET, POST, PUT, DELETE, OPTIONS"}
	//router.Use(cors.New(config))
	//
	//router.POST("/api/authen", controller.Authen)
	//protected := router.Group("/", midldleware.Authentication)
	//protected.GET("/api/setup", controller.Setup)
	////protected.GET("/api/albums/:id", controller.GetAlbumByID)
	//
	//router.Run("localhost:8080")
}
