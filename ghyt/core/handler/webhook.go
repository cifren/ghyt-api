package handler
import (
    "fmt"
    "time"
    "gopkg.in/go-playground/webhooks.v5/github"
    "github.com/kataras/iris"
    "github.com/fatih/structs"
    . "github.com/cifren/ghyt-api/ghyt/core"
    . "github.com/cifren/ghyt-api/ghyt/core/logger"
    . "github.com/cifren/ghyt-api/ghyt/core/job"
)
func GhWebhookHandler(ctx iris.Context, container Container) {
  var errMessage string
  var feedbacks []JobFeedback

	logger := container.Get("logger").(Logger)
  hook, _ := github.New(github.Options.Secret(""))
  payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
		  err := github.ErrEventNotFound
		  errMessage = err.Error()
			fmt.Println(errMessage)
		} else {
		  errMessage = "Error on payload"
			logger.Debug(errMessage)
		}
	} else {
    jobContainer, _ := container.Get("jobContainerFactory").(JobContainerFactory).GetJobContainer(payload)
    logger.Debug(fmt.Sprintf("jobContainer : %+v", jobContainer))

    feedbacks = container.Get("jobRunner").(JobRunner).Run(jobContainer)
    logger.Debug(fmt.Sprintf("Feedbacks : %+v", feedbacks))
  }

  jobLogEntity := JobLog{
    CreatedAt: time.Now(),
    Request: structs.Map(payload),
    ErrorMessage: errMessage,
    Feedback: feedbacks,
  }

  // Record logs
  jobLogRepository, ok := container.Get("jobLogRepository").(JobLogRepositoryInterface)
  if !ok {
    errMessage = fmt.Sprintf(
      "%s, given '%T' does not implement 'JobLogRepositoryInterface'\n",
      "jobLogRepository has to be set in the container, can't reach Job Log repository",
      container.Get("jobLogRepository"),
    )
    fmt.Printf(errMessage)
  } else {
    jobLogRepository.Persist(jobLogEntity)
  }

  ctx.JSON(jobLogEntity)
}
