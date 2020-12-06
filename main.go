package main

import (
	"audit-gateway/conf"
	"audit-gateway/route"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	conf.Init()
	muxRouter := mux.NewRouter()
	route.RegisterRoutes(muxRouter)
	//handler = auth.SsoAuthRequestHandler(handler)
	//handler = logservice.AlilogRequestHandler(handler)
	//http.HandleFunc("/",Serv)
	//userRouter := r.PathPrefix("/user").Subrouter()
	srv := &http.Server{
		Addr:              "0.0.0.0:80",
		Handler:           muxRouter,
		TLSConfig:         nil,
		ReadTimeout:       120 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      120 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	//log.Println(http.ListenAndServe(":8899", handler))
	srv.ListenAndServe()
}
