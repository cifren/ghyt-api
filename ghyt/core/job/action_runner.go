package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/client"
	"fmt"
)

type ActionRunner struct {
	YoutrackClient YoutrackClient
}
func (this ActionRunner) Run(actionConfig Action, jobContainer JobContainer) {
	actionType := ActionRetriever(actionConfig.To, actionConfig.Name)
	client := this.clientResolver(actionConfig.To)

	err := actionType.Run(
		actionConfig,
		jobContainer,
		client,
	)

	if err != nil {
		fmt.Println(err)
	}
}
func (this ActionRunner) clientResolver(clientType string) interface{} {
	var client interface{}

	switch clientType {
		case "youtrack":
			client = this.YoutrackClient
		case "github":
			client = GithubClient{}		
	}

	return client
}
