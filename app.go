package main


import (
  "github.com/kataras/go-template/html"
  "github.com/kataras/iris"
)

type indexpage struct {
    Title   string
    Message string
}
  



func main() {

  iris.UseTemplate(html.New(html.Config{
    Layout: "layout.html",
  })).Directory("./templates", ".html") // the .Directory() is optional also, defaults to ./templates, .html
  // Note for html: this is the default iris' templaet engine, if zero engines added, then the template/html will be used automatically
  // These lines are here to show you how you can change its default configuration

  iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
    ctx.Write(iris.StatusText(iris.StatusInternalServerError)) // Outputs: Internal Server Error
    ctx.SetStatusCode(iris.StatusInternalServerError)          // 500

    ctx.Log("http status: 500 happened!\n")
  })

  iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
    ctx.Write(iris.StatusText(iris.StatusNotFound)) // Outputs: Not Found
    ctx.SetStatusCode(iris.StatusNotFound)          // 404

    ctx.Log("http status: 404 happened!\n")
  })

  iris.Get("/", func(ctx *iris.Context) {
      //ctx.Write("Hi %s", "iris")
      ctx.Render("index.html", indexpage{"My Page title", "Hello world!"}, iris.RenderOptions{"gzip": true})
  })

  // emit the errors to test them
  iris.Get("/500", func(ctx *iris.Context) {
    ctx.EmitError(iris.StatusInternalServerError) // ctx.Panic()
  })

  iris.Get("/404", func(ctx *iris.Context) {
    ctx.EmitError(iris.StatusNotFound) // ctx.NotFound()
  })

  iris.Listen(":8080")
}



func hi(ctx *iris.Context){
  ctx.Write("Hi %s", "iris")
}
