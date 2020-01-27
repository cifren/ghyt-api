package job

import (
	"testing"
	"github.com/stretchr/testify/require"
    "gopkg.in/go-playground/webhooks.v5/github"
)

func TestRun(t *testing.T) {


	testCases := []struct {
        name string
        payload interface{}
        // which tags is we are looking for
        newExpectedValue map[string]string
    }{
        {
            name: "Test1",
            payload: getPullRequestPayload(),
            newExpectedValue: map[string]string {
                "event.pull_request.state": "myState",
                "event.pull_request.title": "myTitle",
            },
        },
    }

	assert := require.New(t)
	for _, tt := range testCases {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			jobContainerFactory := JobContainerFactory{}
			jobContainer, _ := jobContainerFactory.GetJobContainer(tc.payload)
			for key, value := range tc.newExpectedValue {
				assert.Equal(jobContainer.Get(key), value)
			}
		})
	}
}
func getPullRequestPayload() github.PullRequestPayload {
	p := github.PullRequestPayload {}
	p.PullRequest.State = "myState"
	p.PullRequest.Title = "myTitle"

	return p
}
