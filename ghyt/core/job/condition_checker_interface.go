package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/logger"
)

type ConditionCheckerInterface interface {
	Check(conditionConfig Condition, jobContainer *JobContainer, logger Logger) (bool, error)
}



