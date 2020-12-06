package middleware

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	//64位
	cookieStoreAuthKey = "..."
	//AES encrypt key必须是16或者32位
	cookieStoreEncryptKey = "..."
)

var SessionStore *sessions.CookieStore

func init() {
	SessionStore = sessions.NewCookieStore(
		securecookie.GenerateRandomKey(16),
		securecookie.GenerateRandomKey(16),
		//[]byte(cookieStoreAuthKey),
		//[]byte(cookieStoreEncryptKey),
	)

	SessionStore.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60 * 15,
	}

}
