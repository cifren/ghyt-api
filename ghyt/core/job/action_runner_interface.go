package job

import (
	. "github.com/cifren/ghyt-api/ghyt/core/job/tools"
	. "github.com/cifren/ghyt-api/ghyt/core/config"
)


type ActionRunnerInterface interface {
	Run(Action, JobContainer) error
}
