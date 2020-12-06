package service

import (
	"audit-gateway/model"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

func AuditProxy(w http.ResponseWriter, req *http.Request) {
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
				proxy.ServeHTTP(w, req)
				break
			}
			//req.Host = Url.Host
			return
		}
	}
	w.WriteHeader(404)
}
