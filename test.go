package main

import (
	"fmt"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Comment struct {
	Content  string `json:"content" form:"content"`
	VideoId  uint   `json:"videoId" form:"videoId"`
	ParentId uint   `json:"parentId" form:"parentId"`
	FirstId  uint   `json:"firstId" form:"firstId"`
}

func main() {
	m := make([]Comment, 0)
	db, err := gorm.Open("mysql", "root:sakura@tcp(39.107.48.224:3306)/giligili?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

	}
	if err := db.Table("comments").Find(m).Error; err != nil {
		fmt.Println(err)
	}
	err.
		fmt.Println(m)

}
