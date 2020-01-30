package handler
import (
    "fmt"
    "gopkg.in/go-playground/webhooks.v5/github"
    "github.com/kataras/iris"
    . "github.com/cifren/ghyt-api/ghyt/core"
    . "github.com/cifren/ghyt-api/ghyt/core/logger"
    . "github.com/cifren/ghyt-api/ghyt/core/job"
//     "bytes"
	//"net/http"
// 	"encoding/json"
// 	"github.com/cifren/ghyt-api/youtrack/core"
)
func GhWebhookHandler(ctx iris.Context, container Container) {
//
// 	requestBody, _ := json.Marshal(map[string]string{
// 		"name": "nok",
// 	})
//
// 	request := core.Request{
//         //QueryParams: make(url.Values),
//         Endpoint: "issueTags",
//         Body: bytes.NewBuffer(requestBody),
//     }
//     client := core.Client{
//         Url: "https://cospirit.myjetbrains.com/youtrack/api",
//         Token: "Bearer perm:RnJhbmNpc19MZWNvcQ==.NjQtMg==.zwQd0afFB3HJuSONYJfrncHxyXyP9y",
//     }
// 	resp, _ := client.Post(request)
// 	fmt.Printf("%v\n", resp.Status)
// 	//panic("done")
// 	return

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
