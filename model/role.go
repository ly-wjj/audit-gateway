package model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name        string      `gorm:"varchar(100) notnull 'name'"`
	Description string      `gorm:"varchar(255) 'description'"`
	Privileges  []Privilege `gorm:"many2many:role_privilege;"`
}

func (role *Role) GetPrivileges() (privileges []Privilege, err error) {
	err = DB.Model(role).Related(&privileges, "Privileges").Error
	return privileges, err
}
