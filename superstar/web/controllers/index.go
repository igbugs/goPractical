package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"superstar/datasource"
	"superstar/models"
	"superstar/services"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"DataList": dataList,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"Data":  data,
		},
	}
}

// http://localhost:8080/search?country=巴西
func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	dataList := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Data":  dataList,
		},
	}
}

// 集群多服务器的时候，才用得上这个接口
// 性能优化的时候才考虑，加上本机的SQL缓存
// http://localhost:8080/clearcache
func (c *IndexController) GetClearCache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}
	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}
