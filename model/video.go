package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Title  string `json:"Title"`
	Info   string `json:"info"`
	Url    string `json:"url" form:"url"`
	Avatar string `json:"avatar"`
	UserId uint   `json:"userId"`
}
