package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func UploadVideo(c *gin.Context) {
	var s service.UploadTokenService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Post())
	}
}
