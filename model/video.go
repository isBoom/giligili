package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Title string `json:"Title"`
	Info  string `json:"info"`
}
