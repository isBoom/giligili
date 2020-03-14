package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type VideoCommentService struct {
	Content  string `json:"content" form:"content"`
	VideoId  uint   `json:"videoId" form:"videoId"`
	ParentId uint   `json:"parentId" form:"parentId"`
}
type Comments struct {
	ID uint
	VideoCommentService
	Child []Comments `json:"child" form:"child"`
}

func (s *VideoCommentService) Add(user *model.User) serializer.Response {
	if user == nil {
		return serializer.Response{
			Code: 5001,
			Msg:  "未登录，请登陆后再进行评论",
		}
	}
	if s.ParentId != 0 {
		com := model.Comment{}
		if err := model.DB.Where("id = ? and video_id = ?", s.ParentId, s.VideoId).Find(&com).Error; err != nil {
			return serializer.Response{
				Msg: "回复的评论不存在",
			}
		} else if com.UserId == user.ID {
			return serializer.Response{
				Msg: "不能评论自己哟",
			}
		}

	}
	c := model.Comment{
		UserId:   user.ID,
		VideoId:  s.VideoId,
		ParentId: s.ParentId,
		Content:  s.Content,
	}
	if err := model.DB.Create(&c).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "评论失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildComment(c)}
}

func (s *VideoCommentService) Get(c *gin.Context) serializer.Response {
	mods := []Comments{}
	if err := model.DB.Model(&model.Comment{}).Where("video_id = ? and parent_id = 0", c.Param("id")).Find(&mods).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论失败",
			Error: err.Error(),
		}
	}
	for i, mod := range mods {
		temp := []Comments{}
		model.DB.Where("video_id = ? and parent_id = ?", mod.VideoId, mod.ID).Find(&temp)
		mods[i].Child = temp
	}
	return serializer.Response{Data: mods}
}
