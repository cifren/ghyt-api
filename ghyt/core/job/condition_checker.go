package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/logger"
)

type ConditionChecker struct {}
func (this ConditionChecker) Check(
		conditionConfig Condition,
		jobContainer *JobContainer,
		logger Logger,
	) bool {
	conditionType := ConditionRetriever(conditionConfig.Name)

	result, validationErrorMessage := conditionType.Check(
		conditionConfig,
		jobContainer,
	)
	if !result {
		logger.Debug(validationErrorMessage)
	}

	return result
}

