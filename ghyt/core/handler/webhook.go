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
    // "reflect"
    // "github.com/cifren/ghyt/core/event"
    // "github.com/cifren/ghyt/core/action"

)
func GhWebhookHandler(ctx iris.Context, container Container)  {
    // hook, _ := github.New(github.Options.Secret(""))
    // payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
    // if err != nil {
    //  if err == github.ErrEventNotFound {
    //      // ok event wasn't one of the ones asked to be parsed
    //      fmt.Println(github.ErrEventNotFound)
    //  } else {
    //      fmt.Printf("Error on payload: %+v\n", err)
    //  }
    // }
    // fmt.Printf("Payload type: %+v\n", reflect.TypeOf(payload))
    // switch payload.(type) {
    //  case github.PushPayload:
    //      release := payload.(github.PushPayload)
    //      // Do whatever you want from here...
    //      fmt.Printf("%+v\n", release)
    //  case github.PullRequestPayload:
    //      release := payload.(github.PullRequestPayload)

            //event := event.NewPullRequestEvent()
            // event.Variables["id"] = strconv.FormatInt(release.PullRequest.Number, 10)
            // event.Variables["description"] = string(release.PullRequest.Title)
            // myAction := action.AddTag{Name: event.GetVariables()["id"]}
            // runner := action.ActionRunner{}
            // runner.RunIt(myAction)

    //  case github.PingPayload:
    //      release := payload.(github.PingPayload)
    //      // Do whatever you want from here...
    //      fmt.Printf("%+v\n", release)
    //  default:
    //      fmt.Printf("Event without payload : %+v\n", reflect.TypeOf(payload))
    // }
    //event := event.EventManager.GetGhEvent(payload)
    conf := config.GetConf()
    jobContainer := NewJobContainer()
    jobContainer.Set("event.pull_request.state", "open")
    jobContainer.Set("event.pull_request.title", "connect-1517 lol")
    for _, job := range conf {
        runJob(jobContainer, job, container)
    }
}
func runJob(jobContainer *JobContainer, job Job, container Container) {
    logger := container.Get("logger").(Logger)
    logger.Debug(fmt.Sprintf(
        "Conditions found: %x",
        len(job.Conditions),
    ))
    conditionChecker := ConditionChecker{}
    for _, condition := range job.Conditions {
        // varContainer is in ref in case persistName has been set
        if !conditionChecker.Check(condition, jobContainer, logger) {
            logger.Debug(fmt.Sprintf(
                "Condition refused '%s'",
                condition.Name,
            ))
            // quit without executing actions
            return
        }
        logger.Debug(fmt.Sprintf("Condition success '%s'", condition.Name))
    }
    logger.Debug(fmt.Sprintf(
        "Actions found: %x",
        len(job.Actions),
    ))

    actionRunner := container.Get("actionRunner").(ActionRunner)
    // run all actions
    for _, action := range job.Actions {
        logger.Debug(fmt.Sprintf(
            "Run action '%s'",
            action.Name,
        ))
        actionRunner.Run(action, *jobContainer)
    }
}
