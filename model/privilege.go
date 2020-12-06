package model

import "github.com/jinzhu/gorm"

type Privilege struct {
	gorm.Model
	Ptype  string
	Host   string `gorm:"host"`
	Path   string `gorm:"path"`
	Method string `gorm:"method"`
	V3     string `gorm:"v3"`
	V4     string `gorm:"v4"`
	V5     string `gorm:"v5"`
}
