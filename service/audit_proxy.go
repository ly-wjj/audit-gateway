package service

import (
	"github.com/sirupsen/logrus"

	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"audit-gateway/middleware"
	"audit-gateway/model"
)

func CacheService(w http.ResponseWriter, r *http.Request) {
	value, _ := middleware.GetGoCacheData("route_key")
	fmt.Print(value)
	fmt.Fprintf(w, "ss")
}

func AuditProxy(w http.ResponseWriter, req *http.Request) {
	session, _ := middleware.SessionStore.Get(req, sessionCookieName)
	requestLogger := logrus.WithFields(logrus.Fields{
		"host":      req.Host,
		"uri":       req.URL.Path,
		"user":      session.Values["user"],
		"source_ip": req.RemoteAddr,
	})
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		requestLogger.Info("Forbidden")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	var routes []model.Route
	routes, err := model.GetRoutes(req.Host)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	for _, v := range routes {
		r, _ := regexp.Compile(v.Path)
		if v.Path == req.URL.Path || strings.HasPrefix(req.URL.Path, v.Path) || r.MatchString(req.URL.Path) {
			upstream_infos, _ := model.GetUpstreamInfo(v.Upstream)
			for _, upstream := range upstream_infos {
				Url, _ := url.Parse(upstream.UpstreamAddr)
				proxy := httputil.NewSingleHostReverseProxy(Url)
				d := proxy.Director
				proxy.Director = func(r *http.Request) {
					d(r)
					r.Host = Url.Host
					r.URL.Path = req.URL.Path
					if len(v.Path) > 1 && strings.HasSuffix(v.Path, "/") {
						if strings.HasPrefix(req.URL.Path, v.Path) {
							r.URL.Path = strings.Replace(req.URL.Path, v.Path, "/", 1)
						}
					}
				}
				requestLogger.Info("success")
				proxy.ServeHTTP(w, req)
			}
			//req.Host = Url.Host
			return
		}
	}
	w.WriteHeader(404)
}
