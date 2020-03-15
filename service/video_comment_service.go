package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"time"
)

type VideoCommentService struct {
	Content  string `json:"content" form:"content"`
	VideoId  uint   `json:"videoId" form:"videoId"`
	ParentId uint   `json:"parentId" form:"parentId"`
}
type Comments struct {
	ID        uint            `json:"id" form:"id"`
	UserId    uint            `json:"-" form:"-"`
	Users     serializer.User `json:"users" form:"users"`
	CreatedAt time.Time       `json:"createdAt" form:"createdAt"`
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
	users := make([]serializer.User, 0)
	mods := make([]Comments, 0)
	res := make([]Comments, 0)
	mapUser := make(map[uint]*serializer.User)

	//指示最高级评论
	mapFather := make(map[uint]uint)
	//指示最高级评论所在切片下标
	mapIndex := make(map[uint]uint)
	if err := model.DB.Model(&model.Comment{}).Where("video_id = ?", c.Param("id")).Find(&mods).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论失败",
			Error: err.Error(),
		}
	}
	if err := model.DB.Model(&serializer.User{}).Select("DISTINCT users.id,users.user_name,users.nickname,users.avatar").Joins("left JOIN comments on users.id = comments.user_id ").Where("video_id = ?", c.Param("id")).Find(&users).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论失败",
			Error: err.Error(),
		}
	}
	for i, mod := range users {
		mapUser[mod.ID] = &users[i]
	}
	//实现扁平数据二维化
	for _, mod := range mods {
		mod.Users = *mapUser[mod.UserId]
		if mod.ParentId == 0 {
			mapIndex[mod.ID] = uint(len(res))
			res = append(res, mod)
		} else {
			if mapFather[mod.ParentId] == 0 {
				//此时为二级评论
				mapFather[mod.ID] = mod.ParentId
			} else {
				//n级评论
				mapFather[mod.ID] = mapFather[mod.ParentId]
			}
			//最高级评论添加二级到n级子评论
			res[mapIndex[mapFather[mod.ID]]].Child = append(res[mapIndex[mapFather[mod.ID]]].Child, mod)
		}
	}
	return serializer.Response{Data: res}
}
