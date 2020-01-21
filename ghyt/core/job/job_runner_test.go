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
        {
			name: "First condition failed, no action",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-fail-1",
                        },
                        {
                            Name: "condition-success-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                        {
                            Name: "action-noexec-2",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
			},
        },
        {
			name: "Second condition failed, no action",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                        {
                            Name: "condition-fail-1",
                        },
                        {
                            Name: "condition-success-2",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                        {
                            Name: "action-noexec-2",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
			},
		},
        {
			name: "Third condition failed, no action",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                        {
                            Name: "condition-success-2",
                        },
                        {
                            Name: "condition-fail-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                        {
                            Name: "action-noexec-2",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
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
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
			},
        },
        {
			name: "All conditions passed, first action crashed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                        {
                            Name: "condition-success-2",
                        },
                        {
                            Name: "condition-success-3",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-fail-1",
                        },
                        {
                            Name: "action-success-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
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
                        {
							Name: "condition-success-3",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-fail-1",
							ErrorMessage: "Beautiful crash !",
						},
					},
				},
			},
        },
        {
			name: "All conditions passed, second action crashed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                        {
                            Name: "condition-success-2",
                        },
                        {
                            Name: "condition-success-3",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                        {
                            Name: "action-fail-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
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
                        {
							Name: "condition-success-3",
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
							Name: "action-fail-1",
							ErrorMessage: "Beautiful crash !",
						},
					},
				},
			},
        },
        {
			name: "2 jobs, one passed, the other failed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                    },
                },
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-fail-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-success-1",
							ErrorMessage: "",
						},
					},
				},
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
			},
        },
        {
			name: "2 jobs, one failed, the other passed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-fail-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                    },
                },
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-success-1",
							ErrorMessage: "",
						},
					},
				},
			},
        },
        {
			name: "2 jobs, 2 passed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                    },
                },
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-success-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-success-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-success-1",
							ErrorMessage: "",
						},
					},
				},
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-success-1",
                            ErrorMessage: "",
                            Result: true,
                        },
                    },
					ActionFeedbacks: []ActionFeedback {
						{
							Name: "action-success-1",
							ErrorMessage: "",
						},
					},
				},
			},
        },
        {
			name: "2 jobs, 2 failed",
			givenConf: []config.Job {
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-fail-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                    },
                },
                {
                    Conditions: []config.Condition{
                        {
                            Name: "condition-fail-1",
                        },
                    },
                    Actions: []config.Action{
                        {
                            Name: "action-noexec-1",
                        },
                    },
                },
            },
			newExpectedValue: []JobFeedback {
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
				{
                    ConditionFeedbacks: []ConditionFeedback {
                        {
							Name: "condition-fail-1",
                            ErrorMessage: "test case condition failed",
                            Result: false,
                        },
                    },
				},
			},
        },

//         {
//          name: "2 jobs, 2 passed"
//          name: "2 jobs, 2 failed"
// 		   }
    }

	assert := require.New(t)

	for _, tt := range testCases{
		tc := tt

		jobRunner := JobRunner{
			ActionRunner: MockActionRunner{},
			ConditionChecker: MockConditionChecker{},
	        Logger: logger.NewLogger(logger.DEBUG),
	        Configuration: tc.givenConf,
		}

		jobContainer := tools.NewJobContainer()

		actualJobFeedbacks := jobRunner.Run(jobContainer)

		t.Run(tc.name, func(t *testing.T) {
			for expectedJobFdIndex, expectedJobFd := range tc.newExpectedValue {

				assert.Equal(len(expectedJobFd.ConditionFeedbacks), len(actualJobFeedbacks[expectedJobFdIndex].ConditionFeedbacks))
				assert.Equal(len(expectedJobFd.ActionFeedbacks), len(actualJobFeedbacks[expectedJobFdIndex].ActionFeedbacks))

				for conditionFdIndex, conditionFd := range expectedJobFd.ConditionFeedbacks{
					assert.Equal(conditionFd.Name, actualJobFeedbacks[expectedJobFdIndex].ConditionFeedbacks[conditionFdIndex].Name)
					assert.Equal(conditionFd.Result, actualJobFeedbacks[expectedJobFdIndex].ConditionFeedbacks[conditionFdIndex].Result)
					assert.Equal(conditionFd.ErrorMessage, actualJobFeedbacks[expectedJobFdIndex].ConditionFeedbacks[conditionFdIndex].ErrorMessage)
				}

				for actionFdIndex, actionFd := range expectedJobFd.ActionFeedbacks{
					assert.Equal(actionFd.Name, actualJobFeedbacks[expectedJobFdIndex].ActionFeedbacks[actionFdIndex].Name)
					assert.Equal(actionFd.ErrorMessage, actualJobFeedbacks[expectedJobFdIndex].ActionFeedbacks[actionFdIndex].ErrorMessage)
				}
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
			return false, errors.New("test case condition failed")
		default:
			panic(fmt.Sprintf("Case not found, conditionName given : %v", conditionConfig.Name))
	}
}
