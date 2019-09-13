package handler

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
	"github.com/kataras/iris"
	"reflect"
	"github.com/cifren/ghyt/internal/event"
	"github.com/cifren/ghyt/internal/action"
	"strconv"
)

func WebhookHandler(ctx iris.Context)  {
	hook, _ := github.New(github.Options.Secret("plapodwoainjagbwnaodiopONUnad"))
	fmt.Println("/webhook")
	payload, err := hook.Parse(ctx.Request(), github.PingEvent, github.PushEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
			fmt.Println(github.ErrEventNotFound)
		}
	}
	fmt.Println(reflect.TypeOf(payload))
	switch payload.(type) {
		case github.PushPayload:
			release := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)
		case github.PullRequestPayload:
			release := payload.(github.PullRequestPayload)
			
			event := event.NewPullRequestEvent()
			event.Variables["id"] = strconv.FormatInt(release.PullRequest.Number, 10)
			event.Variables["description"] = string(release.PullRequest.Title)

			myAction := action.AddTag{Name: event.GetVariables()["id"]}

			runner := action.ActionRunner{}
			runner.RunIt(myAction)
			
		case github.PingPayload:
			release := payload.(github.PingPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)
		default:
			fmt.Printf("Event without payload : %+v\n", reflect.TypeOf(payload))
	}
}