package main

import (
	"bluenote/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// gin.Default returns an *gin.Engine
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//UserHandler.RegisterRoutes(*gin.Engine) register routes for gin server
	c := &web.UserHandler{}
	c.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Panic(err)
	}
}
