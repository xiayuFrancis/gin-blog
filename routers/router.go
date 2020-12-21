package routers

import (
	"gin-blog/pkg/setting"
	v1 "gin-blog/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags", v1.EditTags)
		apiv1.DELETE("/tags", v1.DeleteTags)
	}

	return r
}
