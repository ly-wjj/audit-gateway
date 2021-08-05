package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := SessionStore.Get(r, sessionCookieName)
		//Do stuff here
		requestLogger := logrus.WithFields(logrus.Fields{
			"host":      r.Host,
			"uri":       r.URL.Path,
			"user":      session.Values["user"],
			"source_ip": r.RemoteAddr,
		})
		requestLogger.Info()
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
