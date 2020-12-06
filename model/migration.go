package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&Route{}, &UpstreamInfo{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Group{})
	DB.AutoMigrate(&Role{}, &Privilege{})
	//InitData()
	//SetPrivilege()
}

func InitData() {
	route := Route{
		Name:     "test",
		Host:     "liya.test.com",
		Path:     "/",
		Upstream: "test_server",
	}
	upstreaminfo := UpstreamInfo{
		Name:         "test_server",
		UpstreamAddr: "https://www.baidu.com",
		Path:         "/",
	}
	privilege := Privilege{
		Ptype:  "allow",
		Host:   "liya.test.com",
		Path:   "/",
		Method: "GET",
	}
	var privle []Privilege
	privle = append(privle, privilege)
	role := Role{
		Name:       "admin",
		Privileges: privle,
	}
	var roles []Role
	roles = append(roles, role)
	group := Group{
		GroupName:   "admin",
		Description: "管理员组",
		Roles:       roles,
	}
	var groups []Group
	groups = append(groups, group)
	user := User{
		UserName: "admin",
		Nickname: "admin",
		Groups:   groups,
	}
	user.SetPassword("test123456")
	DB.Create(&route)
	DB.Create(&upstreaminfo)
	DB.Create(&user)
}
