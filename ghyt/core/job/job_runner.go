package job

import (
	"fmt"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/logger"
	. "github.com/cifren/ghyt/core/job/tools"
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

        // jobContainer is in ref in case persistName has been set
        if !conditionChecker.Check(condition, &jobContainer, this.Logger) {
            conditionFeedback.Result = false
            this.Logger.Debug(fmt.Sprintf(
                "Condition refused '%s'",
                condition.Name,
            ))

            // quit without executing actions
            return
        } else {
            conditionFeedback.Result = true
        }

        jobFeedback.ConditionFeedbacks = append(jobFeedback.ConditionFeedbacks, conditionFeedback)
        this.Logger.Debug(fmt.Sprintf("Condition success '%s'", condition.Name))
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
        this.ActionRunner.Run(action, jobContainer)
        jobFeedback.ActionFeedbacks = append(jobFeedback.ActionFeedbacks, actionFeedback)
    }
}

type JobFeedback struct {
	// Alert if Job crashes
    ErrorMessage string
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
