package main

import (
	"bluenote/internal/repository"
	"bluenote/internal/repository/dao"
	"bluenote/internal/service"
	"bluenote/internal/web"
	"bluenote/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func main() {
	db := InitDB()
	server := InitWebServer()

	// session
	store := cookie.NewStore([]byte("secret"))
	login := middleware.NewLoginMiddlewareBuilder()
	server.Use(sessions.Sessions("sess", store), login.CheckLogin())

	c := initUser(db)

	c.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		log.Panic(err)
	}

}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/bluenote"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func InitWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
	}))
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	c := web.NewUserHandler(svc)
	return c
}
