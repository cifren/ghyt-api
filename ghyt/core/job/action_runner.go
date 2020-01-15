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
func (this ActionRunner) clientResolver(clientTypeName string) interface{} {
	var client interface{}

	switch clientTypeName {
		case "youtrack":
			client = this.YoutrackClient
		case "github":
			client = GithubClient{}
		default:
			panic(fmt.Sprintf("Client type not found, given : %#v", clientTypeName))
	}

	return client
}
