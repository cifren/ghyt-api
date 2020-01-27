package job

import (
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
	. "github.com/cifren/ghyt-api/ghyt/core/logger"
)

type ConditionCheckerInterface interface {
	Check(conditionConfig Condition, jobContainer *JobContainer, logger Logger) (bool, error)
}



