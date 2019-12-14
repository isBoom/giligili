package serializer

import (
	"singo/model"
)

type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
func BuildVideos(item []model.Video) (videos []Video) {
	for _, value := range item {
		video := BuildVideo(value)
		videos = append(videos, video)
	}
	return videos
}
