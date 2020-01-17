package job

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cifren/ghyt/core/logger"
	"github.com/cifren/ghyt/core/job/tools"
	"github.com/cifren/ghyt/core/config"
)

func TestJobRunnerRun(t *testing.T) {

	testCases := []struct {
        name string
        // which tags is we are looking for
        newExpectedValue []JobFeedback
        errExpect bool
    }{
        {
			name: "All conditions and all actions passed",
			newExpectedValue: []JobFeedback {
				{
					ErrorMessage: "",
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "test-equal",
                            ErrorMessage: "",
                            Result: true,
                        },
                        {
							Name: "test-regex",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "test-addTag",
							ErrorMessage: "",
						},
						{
							Name: "test-removeTag",
							ErrorMessage: "",
						},
					},
				},
			},
        },
//         {
// 			name: "All conditions and all actions passed",
// 			name: "First condition failed, no action",
// 			name: "Second condition failed, no action",
// 			name: "Third condition failed, no action",
// 			name: "All conditions passed, first action crashed",
// 			name: "All conditions passed, second action crashed",
// 		}
    }

	assert := require.New(t)

	for index, tt := range testCases{
		tc := tt

		jobRunner := JobRunner{
			ActionRunner: MockActionRunner{},
			ConditionChecker: MockConditionChecker{},
	        Logger: logger.NewLogger(logger.DEBUG),
	        Configuration: GetConf(),
		}

		jobContainer := tools.NewJobContainer()
		jobContainer.Set("Title", "test1")

		jobFeedbacks := jobRunner.Run(jobContainer)

		t.Run(tc.name, func(t *testing.T) {

			actualJobFd := jobFeedbacks[index]
			assert.Equal(tc.newExpectedValue[index].ErrorMessage, actualJobFd.ErrorMessage)
			assert.Equal(len(tc.newExpectedValue[index].ConditionFeedbacks), len(actualJobFd.ConditionFeedbacks))
			assert.Equal(len(tc.newExpectedValue[index].ActionFeedbacks), len(actualJobFd.ActionFeedbacks))

			for conditionFdIndex, conditionFd := range tc.newExpectedValue[index].ConditionFeedbacks{
				assert.Equal(conditionFd.Name, actualJobFd.ConditionFeedbacks[conditionFdIndex].Name)
				assert.Equal(conditionFd.Result, actualJobFd.ConditionFeedbacks[conditionFdIndex].Result)
				assert.Equal(conditionFd.ErrorMessage, actualJobFd.ConditionFeedbacks[conditionFdIndex].ErrorMessage)
			}

			for actionFdIndex, actionFd := range tc.newExpectedValue[index].ActionFeedbacks{
				assert.Equal(actionFd.Name, actualJobFd.ActionFeedbacks[actionFdIndex].Name)
				assert.Equal(actionFd.ErrorMessage, actualJobFd.ActionFeedbacks[actionFdIndex].ErrorMessage)
			}
		})
	}
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
				},
				{
					Name: "test-regex",
				},
			},
			Actions: []config.Action{
				{
		 			Name: "test-addTag",
				},
				{
		 			Name: "test-removeTag",
				},
			},
		},
	}

	return conf
}
