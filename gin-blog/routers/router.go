package routers

import (
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api/v1"
	"gin-blog/routers/api"
	"gin-blog/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 添加标签列表
		apiv1.POST("/tags", v1.AddTag)
		// 更新标签列表
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签列表
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取多个文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定的文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新增文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定的文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除执行的文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
