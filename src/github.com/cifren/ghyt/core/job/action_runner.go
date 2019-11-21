package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/client"
)

type ActionRunner struct {

}
func (this ActionRunner) Run(actionConfig Action, jobContainer JobContainer) {
	actionType := ActionRetriever(actionConfig.To, actionConfig.Name)
	client := this.clientResolver(actionConfig.To)

	actionType.Run(
		actionConfig,
		jobContainer,
		client,
	)
}
func (this ActionRunner) clientResolver(clientType string) interface{} {
	var client interface{}

	switch clientType {
		case "youtrack":
			client = YoutrackClient{}
		case "github":
			client = GithubClient{}		
	}

	return client
}
