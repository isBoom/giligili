package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CreateVideo(c *gin.Context) {

	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := s.Create()
		c.JSON(200, res)
	}
}
func ShowVideo(c *gin.Context) {
	s := service.ShowVideoServics{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := s.Show(c.Param("id"))
		c.JSON(200, res)
	}
}
func ListVideo(c *gin.Context) {
	s := service.ListVideoServics{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := s.List()
		c.JSON(200, res)
	}
}
func UpdateVideo(c *gin.Context) {
	s := service.UpdateVideoService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(5001, ErrorResponse(err))
	} else {
		res := s.Update(c.Param("id"))
		c.JSON(200, res)
	}
}
