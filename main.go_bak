package main

import (
	"github.com/kataras/iris"
	herolib "github.com/kataras/iris/hero"
	. "github.com/cifren/ghyt/core/handler"
	. "github.com/cifren/ghyt/core"
    "path/filepath"
	"runtime"
	// TODO : implement log
	// log "github.com/sirupsen/logrus"
)

func main() {
	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

	def := register()

	webhookHandler := def.Handler(GhWebhookHandler)
	app.Post("/webhook-gh", webhookHandler)

	// http://localhost:8080
	app.Run(iris.Addr(":9001"), iris.WithoutServerError(iris.ErrServerClosed))
}

func getPath() string {
	_, b, _, _ := runtime.Caller(0)
	return  filepath.Dir(b)
}

func register() herolib.Hero {
	def := herolib.New()

	all := make(map[string]interface{})
	all["params"] = params()
	container := Container{All: all}
	container.InitContainer()
	def.Register(container)

	return *def
}

func params() map[string]interface{} {
	return map[string]interface{}{
		"github": map[string]string{
			"github_account": "",
			"github_secret": "",
		},
		"youtrack": map[string]string{
			"url": "",
    		"token": "",
		},
	}
}
