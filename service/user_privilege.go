package service

import (
	"fmt"
	"net/http"
)

func UserPrivileges(w http.ResponseWriter, r *http.Request) {
	username := "admin"
	privileges, _ := GetUserAllPrivileges(username)
	fmt.Fprintln(w, privileges)
}
