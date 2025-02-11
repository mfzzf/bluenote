package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type LoginJWTMiddlewareBuilder struct {
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			// 不需要登录校验
			return
		}
		Authorization := ctx.GetHeader("Authorization")
		if Authorization == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(Authorization, " ")
		token, err := jwt.Parse(segs[1], func(token *jwt.Token) (interface{}, error) {
			return []byte("abW5nQhlwukKm7gx/BfB2w=="), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
