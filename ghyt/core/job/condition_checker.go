package job

import (
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
	. "github.com/cifren/ghyt-api/ghyt/core/logger"
)

type ConditionChecker struct {}
func (this ConditionChecker) Check(
		conditionConfig Condition,
		jobContainer *JobContainer,
		logger Logger,
	) (bool, error) {
	conditionType := ConditionRetriever(conditionConfig.Name)

	result, validationError := conditionType.Check(
		conditionConfig,
		jobContainer,
	)
	if !result {
		logger.Debug(validationError.Error())
		return false, validationError
	}

	return true, nil
}

