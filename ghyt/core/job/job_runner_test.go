package job

import (
	"testing"
	"strings"
	"errors"
	"fmt"

	"github.com/stretchr/testify/require"

	"github.com/cifren/ghyt/core/logger"
	"github.com/cifren/ghyt/core/job/tools"
	"github.com/cifren/ghyt/core/config"
)

func TestJobRunnerRun(t *testing.T) {

	testCases := []struct {
        name string
        givenConf []config.Job
        // which tags is we are looking for
        newExpectedValue []JobFeedback
        errExpect bool
    }{
        {
			name: "All conditions and all actions passed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                        {
                            Name: "condition-success-2",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                        {
                            Name: "action-success-2",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
					ErrorMessage: "",
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                        {
							Name: "condition-success-2",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-success-1",
							ErrorMessage: "",
						},
						{
							Name: "action-success-2",
							ErrorMessage: "",
						},
					},
				},
			},
        },

//         {
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
	        Configuration: tc.givenConf,
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
func (this MockActionRunner) Run(actionConfig config.Action, jobContainer tools.JobContainer) error {
	switch true {
		case strings.Contains(actionConfig.Name, "action-success"):
			return nil
		case strings.Contains(actionConfig.Name, "action-fail"):
			return errors.New("Beautiful crash !")
		default:
			panic(fmt.Sprintf("Case not found, actionName given : %v", actionConfig.Name))
	}
}

type MockConditionChecker struct {}
func (this MockConditionChecker) Check(
		conditionConfig config.Condition,
		jobContainer *tools.JobContainer,
		logger logger.Logger,
	) (bool, error) {
	switch true {
		case strings.Contains(conditionConfig.Name, "condition-success"):
			return true, nil
		case strings.Contains(conditionConfig.Name, "condition-fail"):
			return false, errors.New("Beautiful crash !")
		default:
			panic(fmt.Sprintf("Case not found, conditionName given : %v", conditionConfig.Name))
	}
}
