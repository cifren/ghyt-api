package core

import (
	"github.com/cifren/ghyt/core/logger"
	"github.com/cifren/ghyt/core/client"
	youtrack "github.com/cifren/youtrack/core"
	"github.com/cifren/ghyt/core/job"
)

type Container struct {
	All map[string]interface{}
}
func(this *Container) InitContainer() {
	if this.All == nil {
		this.All = make(map[string]interface{})
	}
	this.All["logger"] = this.getLogger()
	this.All["youtrackClient"] = this.getYoutrackClient()
	this.All["actionRunner"] = this.getActionRunner()
	this.All["jobRunner"] = this.getJobRunner()
	this.All["jobContainerFactory"] = this.getJobContainerFactory()
}
func(this Container) Get(reference string) interface{} {
	return this.All[reference]
}
func(this Container) getLogger() logger.Logger {
	logger := logger.NewLogger(logger.DEBUG)
	return logger
}
func(this Container) getYoutrackClient() client.YoutrackClient {
	params := this.All["params"].(map[string]interface{})
	youtrackParams := params["youtrack"].(map[string]string)
	youtrackUrl := youtrackParams["url"]
	token := youtrackParams["token"]
	clientYt := youtrack.Client{Url: youtrackUrl, Token: token}

	youtrackClient := client.NewYoutrackClient(clientYt)
	return youtrackClient
}
func(this Container) getActionRunner() job.ActionRunner {
	return job.ActionRunner{YoutrackClient: this.getYoutrackClient()}
}
func(this Container) getJobRunner() job.JobRunner {
	return job.ActionRunner{
		ActionRunner: this.getActionRunner(),
		ConditionChecker: ConditionChecker{},
		Logger: this.getLogger(),
	}
}
func(this Container) getJobContainerFactory() job.JobContainerFactory {
	return job.JobContainerFactory{
	}
}
