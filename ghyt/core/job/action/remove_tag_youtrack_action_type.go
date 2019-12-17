package action

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	//. "github.com/cifren/ghyt/core/client"
)

type RemoveTagYoutrackActionType struct {

}
func(this RemoveTagYoutrackActionType) Run(
	actionConfig Action, 
	jobContainer JobContainer,
	clientInterface interface{},
) error {
	// var client YoutrackClient
	// client = clientInterface.(YoutrackClient)

	return nil
}