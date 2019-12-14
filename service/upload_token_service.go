package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"singo/serializer"
)

type UploadTokenService struct {
	FileName string `json:"fileName" form:"fileName"`
}

func (s *UploadTokenService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv(""))
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
		oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + s.FileName

	signedPutRul, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	signedGetRul, err := bucket.SignURL(key, oss.HTTPGet, 600)
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
			"put": signedPutRul,
			"get": signedGetRul,
		},
	}
}
