package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
	//"github.com/cifren/ghyt/internal/model"

	//"reflect"
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	// DB
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

    hook, _ := github.New(github.Options.Secret("plapodwoainjagbwnaodiopONUnad"))
	app.Post("/webhook", func(ctx iris.Context) {
		fmt.Println("/webhook")
		payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent)
        if err != nil {
            if err == github.ErrEventNotFound {
                // ok event wasn't one of the ones asked to be parsed
                fmt.Println(github.ErrEventNotFound)
            }
        }
		//fmt.Printf("/webhook_plop %+v", reflect.TypeOf(payload))
        switch payload.(type) {
            case github.PushPayload:
            	fmt.Println("plop")
                release := payload.(github.PushPayload)
                // Do whatever you want from here...
                fmt.Printf("%+v", release)
            case github.PingPayload:
                release := payload.(github.PingPayload)
                // Do whatever you want from here...
                fmt.Printf("%+v", release)
            case github.PullRequestPayload:
                pullRequest := payload.(github.PullRequestPayload)
                // Do whatever you want from here...
                fmt.Printf("%+v", pullRequest)
            default:
            	fmt.Println("ploppoplop")
				//fmt.Printf("Event without payload : %+v", reflect.TypeOf(payload))
        }
		fmt.Println("done")
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
