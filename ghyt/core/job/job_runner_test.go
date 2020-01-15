package job

import (
	"testing"

	"github.com/cifren/ghyt/core/logger"
	"github.com/cifren/ghyt/core/job/tools"
	"github.com/cifren/ghyt/core/config"
)

func TestJobRunnerRun(t *testing.T) {

	jobRunner := JobRunner{
		ActionRunner: MockActionRunner{},
		ConditionChecker: MockConditionChecker{},
        Logger: logger.NewLogger(logger.DEBUG),
	}

	jobContainer := tools.NewJobContainer()
	jobContainer.Set("Title", "test1")

	jobRunner.Run(GetConf(), jobContainer)
}

type MockActionRunner struct {}
func (this MockActionRunner) Run(actionConfig config.Action, jobContainer tools.JobContainer) {
	// Nothing
}

type MockConditionChecker struct {}
func (this MockConditionChecker) Check(
		conditionConfig config.Condition,
		jobContainer *tools.JobContainer,
		logger logger.Logger,
	) bool {
	return true
}

func GetConf() []config.Job {

	conf := []config.Job {
		{
			Conditions: []config.Condition{
				{
					Name: "test-equal",
					Arguments: map[string]string{
						"variableName": "test-variableName",
						"value": "test-value",
					},
				},
				{
					Name: "test-regex",
					Arguments: map[string]string{
						"variableName": "test-variableName",
						"value": "test-value",
						"persistName": "test-persistName",
					},
				},
			},
			Actions: []config.Action{
				{
					To: "test-youtrack",
		 			Name: "test-addTag",
					Arguments: map[string]string{
						"test-youtrackId": "yt_id",
						"test-tagName": "nok",
					},
				},
				{
					To: "test-youtrack",
		 			Name: "test-removeTag",
					Arguments: map[string]string{
						"youtrackId": "test-yt_id",
						"tagName": "test-nok",
					},
				},
			},
		},
	}

	return conf
}
