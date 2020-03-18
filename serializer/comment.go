package serializer

import (
	"singo/model"
)

type Comment struct {
	User      User
	Id        uint
	VideoId   uint
	FirstId   uint
	ParentId  uint
	Content   string
	CreatedAt int64
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
