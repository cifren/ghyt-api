package job

import (
	"fmt"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
	. "github.com/cifren/ghyt-api/ghyt/core/logger"
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
)

type JobRunner struct {
	ActionRunner ActionRunnerInterface
	ConditionChecker ConditionCheckerInterface
	Logger Logger
	Configuration []Job
}
func (this JobRunner) Run(jobContainer JobContainer) []JobFeedback {
	jobFeedbacks := []JobFeedback{}
	i := 0

    for _, job := range this.Configuration {
        jobFeedback := JobFeedback{}
        this.runJob(jobContainer, job, &jobFeedback)
        jobFeedbacks = append(jobFeedbacks, jobFeedback)
        i++
    }

    return jobFeedbacks
}
func (this JobRunner) runJob(jobContainer JobContainer, job Job, jobFeedback *JobFeedback) {
    this.Logger.Debug(fmt.Sprintf(
        "Conditions found: %x",
        len(job.Conditions),
    ))
    conditionChecker := this.ConditionChecker
    for _, condition := range job.Conditions {
        conditionFeedback := ConditionFeedback{}
        conditionFeedback.Name = condition.Name

		resultCondition, err := conditionChecker.Check(condition, &jobContainer, this.Logger)
        // jobContainer is in ref in case persistName has been set
        if !resultCondition {
            conditionFeedback.Result = false
            conditionFeedback.ErrorMessage = err.Error()
            this.Logger.Debug(fmt.Sprintf(
                "Condition refused '%s' because %s",
                condition.Name,
                err.Error(),
            ))

            // quit without executing actions
        } else {
            conditionFeedback.Result = true
            this.Logger.Debug(fmt.Sprintf("Condition success '%s'", condition.Name))
        }

        jobFeedback.ConditionFeedbacks = append(jobFeedback.ConditionFeedbacks, conditionFeedback)

        if !resultCondition {
            return
        }
    }
    this.Logger.Debug(fmt.Sprintf(
        "Actions found: %x",
        len(job.Actions),
    ))

    // run all actions
    for _, action := range job.Actions {
        actionFeedback := ActionFeedback{}
        actionFeedback.Name = action.Name

        this.Logger.Debug(fmt.Sprintf(
            "Run action '%s'",
            action.Name,
        ))
        err := this.ActionRunner.Run(action, jobContainer)
        if err != nil {
             actionFeedback.ErrorMessage = err.Error()
        }

        jobFeedback.ActionFeedbacks = append(jobFeedback.ActionFeedbacks, actionFeedback)

        if err != nil {
            // stops action loop
            return
        }
    }
}

type JobFeedback struct {
    ConditionFeedbacks []ConditionFeedback
    ActionFeedbacks []ActionFeedback
}
type ConditionFeedback struct {
    Name string
    // Advise if Condition do not pass
    Result bool
	// Alert if Condition crashes
    ErrorMessage string
}
type ActionFeedback struct {
	Name string
	// Alert if Action crashes
	ErrorMessage string
}
