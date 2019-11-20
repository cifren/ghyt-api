package job

import (
	. "github.com/cifren/ghyt/core/config"
)

type ActionTypeInterface interface {
	Run(Action, JobContainer)
}



