package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/export"
	"gin-blog/pkg/qrcode"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrcodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

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
		// 导出标签
		r.POST("/tags/export", v1.ExportTag)
		// 导入标签
		r.POST("tags/import", v1.ImportTag)

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
		// 生成文章的二维码地址
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
