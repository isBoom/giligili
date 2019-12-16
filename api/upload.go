package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/service"
)

func UploadAvatarToken(c *gin.Context) {
	s := service.UploadTokenService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		fmt.Print(s)
		c.JSON(200, s.Post("upload/avatar/"))
	}
}
func UploadVideoToken(c *gin.Context) {
	s := service.UploadTokenService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		fmt.Print(s)
		c.JSON(200, s.Post("upload/video/"))
	}
}
