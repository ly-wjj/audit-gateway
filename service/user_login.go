package service

import (
	"audit-gateway/middleware"
	"audit-gateway/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

var sessionCookieName = "user-session"

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, err := middleware.SessionStore.Get(r, sessionCookieName)
	if err != nil {
		logrus.Error("failed to login")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 登录验证
	name := r.FormValue("username")
	pass := r.FormValue("password")
	var user model.User
	if err := model.DB.Where("user_name = ?", name).First(&user).Error; err != nil {
		log.Fatal("账号或密码错误", nil)
		fmt.Fprintln(w, "账号或密码错误")
		return
	}
	if user.CheckPassword(pass) == false {
		log.Fatal("账号或密码错误", nil)
		fmt.Fprintln(w, "账号或密码错误")
		return
	}
	// 在session中标记用户已经通过登录验证
	session.Values["authenticated"] = true
	session.Values["user"] = name
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "登录成功!", err)
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionStore.Get(r, sessionCookieName)
	auth := session.Values["authenticated"]
	if auth == true {
		user := session.Values["user"]
		fmt.Fprintf(w, "%s", user)
	} else {
		fmt.Fprint(w, "未登录认证")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionStore.Get(r, sessionCookieName)
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionStore.Get(r, sessionCookieName)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "已经登录了")
}
