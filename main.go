package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/cifren/ghyt/core/handler"
)

func main() {
	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

	webhookHandler := hero.Handler(handler.GhWebhookHandler)
	app.Post("/webhook-gh", webhookHandler)

	// http://localhost:8080
	app.Run(iris.Addr(":9001"), iris.WithoutServerError(iris.ErrServerClosed))
}
