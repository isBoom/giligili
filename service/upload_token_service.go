package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime"
	"os"
	"path/filepath"
	"singo/serializer"
)

type UploadTokenService struct {
	FileName string `json:"fileName" form:"fileName"`
}

func (s *UploadTokenService) Post(src string) serializer.Response {
	client, err := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKER"))
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(filepath.Ext(s.FileName))),
	}

	key := src + uuid.Must(uuid.NewRandom()).String() + "_" + s.FileName

	signedPutUrl, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	signedGetUrl, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutUrl,
			"get": signedGetUrl,
		},
	}
}
