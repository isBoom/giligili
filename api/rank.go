package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func DailyRank(c *gin.Context) {
	s := service.DailyRankService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
