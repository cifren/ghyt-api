package job

import (
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
)

type ConditionTypeInterface interface {
	Check(Condition, *JobContainer) (bool, error)
}

