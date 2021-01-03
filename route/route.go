package route

import (
	"audit-gateway/service"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", service.AuditProxy)
	r.HandleFunc("/cache", service.CacheService)
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", service.Login).Methods("POST").Host("liya.test.com")
	userRouter.HandleFunc("/secret", service.Secret).Host("liya.test.com")
	userRouter.HandleFunc("/logout", service.Logout).Host("liya.test.com")
	userRouter.HandleFunc("/privilege", service.UserPrivileges).Host("liya.test.com")

}
