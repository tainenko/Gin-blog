package routers

import (
	"github.com/gin-blog/pkg/setting"
	"github.com/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1"){
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags", v1.PutTags)
		apiv1.DELETE("/tags", v1.DelTags)
	}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
