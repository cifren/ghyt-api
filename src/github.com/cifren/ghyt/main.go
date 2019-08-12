package main

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/cifren/ghyt/internal/model"

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

    fmt.Println("Touch my tralala main")
	// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
	    fmt.Println("Touch my tralala index")
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

    hook, _ := github.New(github.Options.Secret("plapodwoainjagbwnaodiopONUnad"))
	app.Post("/webhook", func(ctx iris.Context) {
	    fmt.Println("Touch my tralala webhook")
	    fmt.Println(ctx.Request())
		payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PullRequestEvent)
        if err != nil {
            if err == github.ErrEventNotFound {
                // ok event wasn't one of the ones asked to be parsed
                fmt.Println(github.ErrEventNotFound)
            }
        }
        switch payload.(type) {

            case github.PingPayload:
                release := payload.(github.PingPayload)
                // Do whatever you want from here...
                fmt.Printf("%+v", release)

            case github.PullRequestPayload:
                pullRequest := payload.(github.PullRequestPayload)
                // Do whatever you want from here...
                fmt.Printf("%+v", pullRequest)
            default:
                fmt.Println(payload)
        }
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
