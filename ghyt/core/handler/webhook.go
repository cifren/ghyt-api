package handler
import (
    "fmt"
    _ "gopkg.in/go-playground/webhooks.v5/github"
    "github.com/kataras/iris"
    . "github.com/cifren/ghyt/core"
    . "github.com/cifren/ghyt/core/logger"
    . "github.com/cifren/ghyt/core/job"
    . "github.com/cifren/ghyt/core/job/tools"
    . "github.com/cifren/ghyt/core/config"
    "github.com/cifren/ghyt/config"
    "github.com/cifren/ghyt/core/event"
)
func GhWebhookHandler(ctx iris.Context, container Container)  {
    hook, _ := github.New(github.Options.Secret(""))
    payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
			fmt.Println(github.ErrEventNotFound)
		} else {
			fmt.Printf("Error on payload: %+v\n", err)
		}
	}

    jobContainer = container.Get("jobContainerFactory").GetJobContainer(payload)

    jobRunner := container.Get("jobRunner").Run(conf, jobContainer)
}
