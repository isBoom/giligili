package serializer

import "singo/model"

type Comment struct {
	User     User
	Id       uint
	VideoId  uint
	ParentId uint
	Content  string
}

func BuildComment(i model.Comment) Comment {
	user, _ := model.GetUser(i.UserId)
	return Comment{
		User:     BuildUser(user),
		Id:       i.ID,
		VideoId:  i.VideoId,
		ParentId: i.ParentId,
		Content:  i.Content,
	}
}
