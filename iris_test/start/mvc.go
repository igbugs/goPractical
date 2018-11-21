package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))

	app.Run(iris.Addr(":8080"))
}

type ExampleController struct {}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType:  "text/html",
		Text: "<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{}{
	return map[string]string{
		"message": "Hello Iris!",
	}
}
