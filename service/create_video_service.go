package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 上传视频
type CreateVideoService struct {
	Title  string `json:"title" form:"title" binding:"required,min=2,max=50"`
	Info   string `json:"info" form:"info" binding:"min=2,max=500"`
	Url    string `json:"url" form:"url" binding:"required`
	Avatar string `json:"avatar" form:"avatar"`
	UserId uint   `json:"userId" form:"userId"`
}

func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {

	v := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		Url:    service.Url,
		Avatar: service.Avatar,
	}

	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			v.UserId = u.ID
		} else {
			return serializer.Response{
				Code: 5001,
				Msg:  "未登录，请登陆后再上传视频",
			}
		}
	}

	if err := model.DB.Create(&v).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(v)}
}
