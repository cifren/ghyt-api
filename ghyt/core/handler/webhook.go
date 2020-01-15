package handler
import (
    "fmt"
    "gopkg.in/go-playground/webhooks.v5/github"
    "github.com/kataras/iris"
    . "github.com/cifren/ghyt/core"
    . "github.com/cifren/ghyt/core/job"
    . "github.com/cifren/ghyt/core/config"
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

    jobContainerFactory := container.Get("jobContainerFactory").(JobContainerFactory)
    jobContainer, _ := jobContainerFactory.GetJobContainer(payload)

    jobRunner := container.Get("jobRunner").(JobRunner).Run([]Job{}, jobContainer)
    fmt.Printf("%v", jobRunner)
}
