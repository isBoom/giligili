package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 上传视频
type CreateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"required,min=0,max=300"`
}

func (service *CreateVideoService) Create() serializer.Response {

	v := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	if err := model.DB.Create(&v).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	fmt.Println(v)
	return serializer.Response{Data: serializer.BuildVideo(v)}
}
