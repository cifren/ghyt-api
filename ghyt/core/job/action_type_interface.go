package job

import (
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
)

type ActionTypeInterface interface {
	Run(Action, JobContainer, interface{}) error
}



