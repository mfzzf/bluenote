package main

import (
	"bluenote/internal/repository"
	"bluenote/internal/repository/dao"
	"bluenote/internal/service"
	"bluenote/internal/web"
	"bluenote/internal/web/middleware"
	ratelimit "bluenote/pkg/ginx/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redisSess "github.com/gin-contrib/sessions/redis"
	"github.com/redis/go-redis/v9"

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

	c := initUser(db)

	c.RegisterRoutes(server)

	//server := gin.Default()
	//server.GET("/hello", func(ctx *gin.Context) {
	//	ctx.String(200, "hello")
	//})

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
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.31.97:3000", "http://192.168.0.2:3000"}, // Removed trailing slash
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "localhost") {
				return true
			}
			if strings.Contains(origin, "192") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
	}))

	store, err := redisSess.NewStore(16, "tcp", "localhost:6379", "", []byte("abW5nQhlwukKm7gx/BfB2w=="), []byte("ZaQqleZrLOznZnKsZdB5FQ=="))
	if err != nil {
		log.Panic(err)
	}
	// check login
	login := middleware.NewLoginJWTMiddlewareBuilder()
	server.Use(sessions.Sessions("sess", store), login.CheckLogin())
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	c := web.NewUserHandler(svc)
	return c
}
