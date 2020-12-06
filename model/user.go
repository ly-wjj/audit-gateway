package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string  `gorm:"size:1000"`
	Roles          []Role  `gorm:"many2many:user_role;"`
	Groups         []Group `gorm:"many2many:user_group;"`
}

type Group struct {
	gorm.Model
	GroupName   string
	Description string
	Roles       []Role `gorm:"many2many:group_role;"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID uint) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := DB.Where("user_name = ?", username).First(&user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) GetGroups() (groups []Group, err error) {
	err = DB.Model(user).Related(&groups, "Groups").Error
	return groups, err
}

func (user *User) GetRoles() (roles []Role, err error) {
	err = DB.Model(user).Related(&roles, "Roles").Error
	return roles, err
}

func (user *User) GetRolePrivileges() (privileges []Privilege, err error) {
	var roles []Role
	err = DB.Model(user).Related(&roles, "Roles").Error
	var privs []Privilege
	for _, role := range roles {
		privs, err = role.GetPrivileges()
		if err == nil {
			privileges = append(privileges, privs...)
		}
	}
	return privileges, err
}

func (group *Group) GetGroupRoles() (roles []Role, err error) {
	err = DB.Model(group).Related(&roles, "Roles").Error
	return roles, err
}

func (group *Group) GetGroupRolePrivileges() (privileges []Privilege, err error) {
	var roles []Role
	err = DB.Model(group).Related(&roles, "Roles").Error
	var privs []Privilege
	for _, role := range roles {
		privs, err = role.GetPrivileges()
		if err == nil {
			privileges = append(privileges, privs...)
		}
	}
	return privileges, err
}
