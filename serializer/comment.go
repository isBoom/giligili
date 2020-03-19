package serializer

import (
	"singo/model"
)

type Comment struct {
	User      User   `json:"user" form:"user"`
	Id        uint   `json:"id" form:"id"`
	VideoId   uint   `json:"videoId" form:"videoId"`
	FirstId   uint   `json:"firstId" form:"firstId"`
	ParentId  uint   `json:"parentId" form:"parentId"`
	Content   string `json:"content" form:"content"`
	CreatedAt int64  `json:"createdAt" form:"createdAt"`
}

func BuildComment(i model.Comment) Comment {
	user, _ := model.GetUser(i.UserId)
	return Comment{
		User:      BuildUser(user),
		Id:        i.ID,
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
