package service

import (
	"audit-gateway/model"
	"fmt"
	"regexp"
	"strings"
)

func GetUserAllPrivileges(username string) (privileges []model.Privilege, err error) {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		fmt.Printf("没有查到相关的数据,%d", username)
	}
	pris, _ := user.GetRolePrivileges()
	privileges = append(privileges, pris...)
	groups, _ := user.GetGroups()
	for _, gro := range groups {
		pris, _ := gro.GetGroupRolePrivileges()
		privileges = append(privileges, pris...)
	}
	return privileges, err
}

func CheckRoute(username string, http_host string, path string, method string) bool {
	privileges, _ := GetUserAllPrivileges(username)
	for _, pri := range privileges {
		if pri.Method == method && pri.Host == http_host {
			r, _ := regexp.Compile(pri.Path)
			if pri.Path == path || strings.HasPrefix(path, pri.Path) || r.MatchString(path) {
				if pri.Ptype == "allow" {
					return true
				}
			}
		}
	}
	return false
}

//func CheckRoute(username string, http_host string, path string) (string, bool) {
//	var routes = make([]model.Route, 0)
//	routes,err := model.GetRoutes(http_host)
//	if err != nil {
//		fmt.Println(err.Error())
//		return "", false
//	}
//	if len(routes) == 0 {
//		if err := model.DB.Find(&routes); err != nil {
//			fmt.Println(err.Error)
//			return "", false
//		}
//		var head = `<!DOCTYPE html><html lang="zh-CN"><head><meta charset="utf-8"></head><body><ul>`
//		for _, v := range routes {
//			head += fmt.Sprintf(`<li>%s: <a href="http://%s%s" target="_blank">http://%s%s</a></li>`, v.Category, v.Host, v.Path, v.Host, v.Path)
//		}
//		head += "</ul></body></html>"
//		//w.Write([]byte(head))
//		return head, false
//	}
//	sort.Slice(routes, func(i, j int) bool {
//		return len(routes[i].Path) > len(routes[j].Path)
//	})
//	for _, v := range routes {
//		r, _ := regexp.Compile(v.Path)
//		if (v.Path == path || strings.HasPrefix(path, v.Path) || r.MatchString(path)) && func() bool {
//			if v.Allow == "all" || v.Allow == "" {
//				return true
//			}
//			for _, o := range strings.Split(v.Allow, ",") {
//				if o == nickname {
//					return true
//				}
//			}
//			return false
//		}() {
//			return "", true
//		}
//		continue
//	}
//	return "", false
//}
