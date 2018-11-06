package controllers

import (
	"github.com/kataras/iris"
	"superstar/services"
	"github.com/kataras/iris/mvc"
	"superstar/models"
	"log"
	"time"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
		dataList := c.Service.GetAll()
		return mvc.View{
			Name: "admin/index.html",
			Data: iris.Map{
				"Title": "管理后台",
				"DataList": dataList,
			},
			Layout: "admin/layout.html",
		}
}

func (c *AdminController) GetEdit() mvc.Result {
	var data *models.StarInfo
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		data = &models.StarInfo{Id: 0}
		log.Fatal(err)
	}
	data = c.Service.Get(id)
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}

	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{
			"name_zh",
			"name_en",
			"avator",
			"birthday",
			"height",
			"weight",
			"club",
			"jersy",
			"country",
			"moreinfo",
		})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path: "/amdin/",
	}
}
