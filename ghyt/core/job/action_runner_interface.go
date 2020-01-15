package job

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
)


type ActionRunnerInterface interface {
	Run(Action, JobContainer)
}
