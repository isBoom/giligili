package serializer

import (
	"singo/model"
)

type Comment struct {
	Id        uint   `json:"id" form:"id"`
	UserName  string `json:"user_name" form:"userName"`
	Nickname  string `json:"nickname" form:"nickname"`
	Avatar    string `json:"avatar" form:"avatar"`
	VideoId   uint   `json:"videoId" form:"videoId"`
	FirstId   uint   `json:"firstId" form:"firstId"`
	ParentId  uint   `json:"parentId" form:"parentId"`
	Content   string `json:"content" form:"content"`
	CreatedAt int64  `json:"createdAt" form:"createdAt"`
}

func BuildComment(i model.Comment) Comment {
	user, _ := model.GetUser(i.UserId)
	return Comment{
		Id:        i.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		VideoId:   i.VideoId,
		ParentId:  i.ParentId,
		Content:   i.Content,
		FirstId:   i.FirstId,
		CreatedAt: i.CreatedAt.Unix(),
	}
}

//func BuildComments(items []model.Comment) (Comments []Comment) {
//	for _, item := range items {
//		comment := BuildComment(item)
//		Comments = append(Comments, comment)
//	}
//	return Comments
//}
