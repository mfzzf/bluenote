package web

import (
	"bluenote/internal/domain"
	"bluenote/internal/service"
	"errors"
	"fmt"
	"net/http"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ErrUserDuplicateEmail = service.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = service.ErrInvalidUserOrPassword

type UserHandler struct {
	Service    *service.UserService
	EmailRe    *regexp.Regexp
	PasswordRe *regexp.Regexp
}

func NewUserHandler(Service *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	EmailRe := regexp.MustCompile(emailRegexPattern, 0)
	PasswordRe := regexp.MustCompile(passwordRegexPattern, 0)

	return &UserHandler{
		Service:    Service,
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
	sess := sessions.Default(ctx)
	sessVal := sess.Get("userId")
	if sessVal != nil {
		user, err := h.Service.Profile(ctx, domain.User{
			ID: sessVal.(int64),
		})
		if err != nil {
			ctx.String(200, "系统错误")
			return
		}
		birthday := time.Unix(user.Birthday, 0).Format("2006-01-02")
		ctx.JSON(200, gin.H{
			"Nickname": user.Nickname,
			"Email":    user.Email,
			"Phone":    user.Phone,
			"Birthday": birthday,
			"AboutMe":  user.AboutMe,
		})
	}
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
	err = h.Service.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, ErrUserDuplicateEmail) {
		ctx.String(200, "邮箱冲突")
	}
	if err != nil {
		//ctx.String(200, "系统异常")
		return
	}
	ctx.String(200, "注册成功")
	fmt.Printf("%v\n", req)

}

func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	user, err := h.Service.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, ErrInvalidUserOrPassword) {
		ctx.String(200, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(200, "系统错误")
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("userId", user.ID)
	err = sess.Save()
	if err != nil {
		ctx.String(200, "系统错误")
		return
	}
	ctx.String(200, "登录成功")
	return
}

// {nickname: "1", birthday: "2025-02-09", aboutMe: "1"}
func (h *UserHandler) Edit(ctx *gin.Context) {
	type editReq struct {
		Nickname string `json:"nickname"`
		Birthday string `json:"birthday"`
		AboutMe  string `json:"aboutMe"`
	}
	var req editReq
	ctx.Bind(&req)

	sess := sessions.Default(ctx)
	sessVal := sess.Get("userId")

	parse, _ := time.Parse("2006-01-02", req.Birthday)
	birthday := parse.Unix()
	err := h.Service.Edit(ctx, domain.User{
		ID:       sessVal.(int64),
		Nickname: req.Nickname,
		Birthday: birthday,
		AboutMe:  req.AboutMe,
	})
	if err != nil {
		ctx.String(200, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "修改成功"})
}
