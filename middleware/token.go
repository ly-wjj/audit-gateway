package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

//自定义一个字符串
var jwtkey = []byte("*.audit.test.com")
var sessionCookieName = os.Getenv("SessionCookieName")

type Claims struct {
	UserName string
	jwt.StandardClaims
}

func CreateToken(userName string) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user/login" {
			next.ServeHTTP(w, r)
			return
		}
		session, err := SessionStore.Get(r, sessionCookieName)
		if err != nil {
			logrus.Error("failed to login")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tokenString := session.Values["Authorization"].(string)
		//vcalidate token formate
		if tokenString == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		logrus.Info(claims.UserName)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
