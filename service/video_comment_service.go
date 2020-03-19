package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"time"
)

type VideoCommentService struct {
	UserId   uint   `json:"userId" form:"userId"`
	Content  string `json:"content" form:"content"`
	VideoId  uint   `json:"videoId" form:"videoId"`
	ParentId uint   `json:"parentId" form:"parentId"`
	FirstId  uint   `json:"firstId" form:"firstId"`
}
type Comments struct {
	ID             uint      `json:"id" form:"id"`
	UserName       string    `json:"user_name" form:"userName"`
	Nickname       string    `json:"nickname" form:"nickname"`
	Avatar         string    `json:"avatar" form:"avatar"`
	CreatedAt      time.Time `json:"-" form:"-"`
	CreatedAtInt64 int64     `json:"createdAt" form:"createdAt"`
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
		if err := model.DB.Where("id = ? and video_id = ?", s.ParentId, s.VideoId).First(&com).Error; err != nil {
			return serializer.Response{
				Msg: "回复的评论不存在",
			}
		}
		if com.ParentId == 0 {
			s.FirstId = com.ID
		} else {
			s.FirstId = com.FirstId
		}
	}

	c := model.Comment{
		UserId:   user.ID,
		VideoId:  s.VideoId,
		ParentId: s.ParentId,
		Content:  s.Content,
		FirstId:  s.FirstId,
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
	//一次性查询所有评论
	mods := make([]Comments, 0)
	//最终返回结果
	res := make([]Comments, 0)
	if err := model.DB.Model(&model.Comment{}).
		Select("comments.id,comments.created_at,content,video_id,first_id,parent_id,user_id,avatar,nickname,user_name").
		Joins("left JOIN users on users.id = comments.user_id ").
		Where("video_id = ? and parent_id = 0", c.Param("id")).
		Order("comments.id desc").Find(&res).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论失败",
			Error: err.Error(),
		}
	}
	if err := model.DB.Model(&model.Comment{}).
		Select("comments.id,comments.created_at,content,video_id,first_id,parent_id,user_id,avatar,nickname,user_name").
		Joins("left JOIN users on users.id = comments.user_id ").
		Where("video_id = ? and parent_id != 0", c.Param("id")).
		Order("comments.id desc").Find(&mods).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论失败",
			Error: err.Error(),
		}
	}
	//标识父评论下标
	mapIndex := make(map[uint]uint)
	for i, mod := range res {
		res[i].CreatedAtInt64 = res[i].CreatedAt.Unix()
		mapIndex[mod.ID] = uint(i)
	}
	//实现扁平数据二维化
	for _, mod := range mods {
		mod.CreatedAtInt64 = mod.CreatedAt.Unix()
		if res[mapIndex[mod.FirstId]].Child == nil {
			res[mapIndex[mod.FirstId]].Child = make([]Comments, 0)
		}
		res[mapIndex[mod.FirstId]].Child = append(res[mapIndex[mod.FirstId]].Child, mod)
		fmt.Println(res)
	}
	return serializer.Response{Data: res}
	//users := make([]serializer.User, 0)
	//mods := make([]Comments, 0)
	//res := make([]Comments, 0)
	//mapUser := make(map[uint]*serializer.User)
	//
	////指示最高级评论
	//mapFather := make(map[uint]uint)
	////指示最高级评论所在切片下标
	//mapIndex := make(map[uint]uint)
	//if err := model.DB.Model(&model.Comment{}).Where("video_id = ?", c.Param("id")).Find(&mods).Error; err != nil {
	//	return serializer.Response{
	//		Code:  5001,
	//		Msg:   "获取评论失败",
	//		Error: err.Error(),
	//	}
	//}
	//if err := model.DB.Model(&serializer.User{}).Select("DISTINCT users.id,users.user_name,users.nickname,users.avatar").Joins("left JOIN comments on users.id = comments.user_id ").Where("video_id = ?", c.Param("id")).Find(&users).Error; err != nil {
	//	return serializer.Response{
	//		Code:  5001,
	//		Msg:   "获取评论失败",
	//		Error: err.Error(),
	//	}
	//}
	//for i, mod := range users {
	//	mapUser[mod.ID] = &users[i]
	//}
	////实现扁平数据二维化
	//for _, mod := range mods {
	//	mod.Users = *mapUser[mod.UserId]
	//	if mod.ParentId == 0 {
	//		mapIndex[mod.ID] = uint(len(res))
	//		res = append(res, mod)
	//	} else {
	//		if mapFather[mod.ParentId] == 0 {
	//			//此时为二级评论
	//			mapFather[mod.ID] = mod.ParentId
	//		} else {
	//			//n级评论
	//			mapFather[mod.ID] = mapFather[mod.ParentId]
	//		}
	//		//最高级评论添加二级到n级子评论
	//		res[mapIndex[mapFather[mod.ID]]].Child = append(res[mapIndex[mapFather[mod.ID]]].Child, mod)
	//	}
	//}
	//return serializer.Response{Data: res}
}
