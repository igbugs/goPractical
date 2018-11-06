package main

import "github.com/kataras/iris"

func main()  {
	app := iris.New()
	// 注册模板
	htmlEngine := iris.HTML("./", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world! --from iris")

	})

	// 填充模板数据，设定使用的模板
	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Titile", "测试页面")
		ctx.ViewData("Content", "hello world! -- from template")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
