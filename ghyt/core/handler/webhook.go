package handler
import (
    "fmt"
    "gopkg.in/go-playground/webhooks.v5/github"
    "github.com/kataras/iris"
    . "github.com/cifren/ghyt-api/ghyt/core"
    . "github.com/cifren/ghyt-api/ghyt/core/logger"
    . "github.com/cifren/ghyt-api/ghyt/core/job"
)
func GhWebhookHandler(ctx iris.Context, container Container) {
	logger := container.Get("logger").(Logger)
    hook, _ := github.New(github.Options.Secret(""))
    payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
			fmt.Println(github.ErrEventNotFound)
		} else {
			logger.Debug("Error on payload: %+v")
		}
	}

  jobContainer, _ := container.Get("jobContainerFactory").(JobContainerFactory).GetJobContainer(payload)
  logger.Debug(fmt.Sprintf("jobContainer : %+v", jobContainer))

  jobRunner := container.Get("jobRunner").(JobRunner).Run(jobContainer)
  logger.Debug(fmt.Sprintf("Feedbacks : %+v", jobRunner))
}
