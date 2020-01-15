package job

import (
    "reflect"
    "fmt"
    "errors"

    "gopkg.in/go-playground/webhooks.v5/github"

	. "github.com/cifren/ghyt/core/job/tools"
)

type JobContainerFactory struct {
}
func (this JobContainerFactory) GetJobContainer(payload interface{}) (JobContainer, error) {
	jobContainer := JobContainer{}

    switch payload.(type) {
		case github.PushPayload:
			release := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)
		case github.PullRequestPayload:
			release := payload.(github.PullRequestPayload)

		    jobContainer = NewJobContainer()
		    jobContainer.Set("event.pull_request.state", release.PullRequest.State)
		    jobContainer.Set("event.pull_request.title", release.PullRequest.Title)
		case github.PingPayload:
			release := payload.(github.PingPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)
		default:
			err := errors.New(fmt.Sprintf(
                "Payload type not recognised, given : %s",
                reflect.TypeOf(payload),
            ))
			return JobContainer{}, err
    }

    return jobContainer, nil
}
