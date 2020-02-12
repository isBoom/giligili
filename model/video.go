package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
)

type Video struct {
	gorm.Model
	Title  string `json:"Title"`
	Info   string `json:"info"`
	Url    string `json:"url" form:"url"`
	Avatar string `json:"avatar"`
	UserId uint   `json:"userId"`
}

func (video *Video) AvatarUrl() string {
	client, _ := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKER"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 60)
	if strings.Contains(signedGetURL, os.Getenv("OSS_UserInfoUrl")+"?Exp") || (video.Avatar == "") {
		signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	}
	return signedGetURL
	//if video.Avatar != "" {
	//	return os.Getenv("OSS_UserInfoUrl") + video.Avatar
	//} else {
	//	return ""
	//}
}

func (video *Video) VideoUrl() string {
	client, _ := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKER"))
	signedGetURL, _ := bucket.SignURL(video.Url, oss.HTTPGet, 3600)
	return signedGetURL
}
