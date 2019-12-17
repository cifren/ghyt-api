package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
)

type ActionTypeInterface interface {
	Run(Action, JobContainer, interface{}) error
}



