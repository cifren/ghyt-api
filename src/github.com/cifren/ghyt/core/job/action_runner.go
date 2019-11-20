package job

import (
	. "github.com/cifren/ghyt/core/config"
)

type ActionRunner struct {

}
func (this ActionRunner) Run(actionConfig Action, jobContainer JobContainer) {
	actionType := ActionRetriever(actionConfig.To, actionConfig.Name)

	actionType.Run(
		actionConfig,
		jobContainer,
	)
}
