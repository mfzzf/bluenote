package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	EmailRe    *regexp.Regexp
	PasswordRe *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	EmailRe := regexp.MustCompile(emailRegexPattern, 0)
	PasswordRe := regexp.MustCompile(passwordRegexPattern, 0)

	return &UserHandler{
		EmailRe:    EmailRe,
		PasswordRe: PasswordRe,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/users")
	userGroup.GET("/profile", h.Profile)
	userGroup.POST("/signup", h.SignUp)
	userGroup.POST("/login", h.Login)
	userGroup.POST("/edit", h.Edit)
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(200, "user profile")
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	//构建创建用户结构体，前端传来Email,Password和确认密码
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	//根据content-type 解析数据 json/form
	var req SignUpReq
	err := ctx.Bind(&req)
	if err != nil {
		return
	}

	if isMatch, _ := h.EmailRe.MatchString(req.Email); !isMatch {
		ctx.String(http.StatusOK, "你的邮箱格式错误")
		fmt.Printf("%v", req)
		return
	}

	if isMatch, _ := h.PasswordRe.MatchString(req.Password); !isMatch {
		ctx.String(http.StatusOK, "你的密码强度太低")
		fmt.Printf("%v", req)
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(200, "你输入的密码不一致")
		return
	}

	//写入数据库
	ctx.String(200, "注册成功")
	fmt.Printf("%v\n", req)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	ctx.String(200, "login")
}

func (h *UserHandler) Edit(ctx *gin.Context) {
	ctx.String(200, "edit")
}
