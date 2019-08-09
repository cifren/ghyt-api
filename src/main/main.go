package main

import (
    "github.com/kataras/iris"
)

func main() {

    app := iris.New()

    // Method:   GET
    // Resource: http://localhost:8080
    app.Handle("GET", "/", func(ctx iris.Context) {
        ctx.HTML("<h1>Welcome</h1>")
    })

    // http://localhost:8080
    // http://localhost:8080/ping
    // http://localhost:8080/hello
    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}