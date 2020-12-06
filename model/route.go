package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	Name       string
	Host       string `gorm:"varchar(255) notnull 'host'"`
	Path       string `gorm:"varchar(255) notnull 'path'"`
	Upstream   string `gorm:"varchar(255) notnull 'upstream'"`
	PriorityId uint   `gorm:"default:100"`
	Category   string `gorm:"varchar(1024) 'category'"`
}

type UpstreamInfo struct {
	gorm.Model
	Name         string
	UpstreamAddr string
	Path         string
	Category     string
}

func (upstream *UpstreamInfo) SetUpstreamInfo(upstream_addr string, path string, category string) {
	upstream.UpstreamAddr = upstream_addr
	upstream.Path = path
	upstream.Category = category

}

func GetUpstreamInfo(upstream_name string) (ups []UpstreamInfo, err error) {
	var upstream UpstreamInfo
	result := DB.Model(&upstream).Where("name = ?", upstream_name).Find(&ups)
	return ups, result.Error
}

func GetRoutes(host string) (routes []Route, err error) {
	var route Route
	result := DB.Model(&route).Where("host = ?", host).Order("priority_id desc").Find(&routes)
	return routes, result.Error
}
