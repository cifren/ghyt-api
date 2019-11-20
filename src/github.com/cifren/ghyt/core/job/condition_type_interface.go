package job

import (
	. "github.com/cifren/ghyt/core/config"
)

type ConditionTypeInterface interface {
	Check(Condition, JobContainer) (bool, string)
}

