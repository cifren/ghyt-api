package action

import (
	. "github.com/cifren/ghyt/core/job/tools"
	. "github.com/cifren/ghyt/core/config"
	. "github.com/cifren/ghyt/core/client"
	. "github.com/cifren/youtrack"
	"errors"
	"fmt"
)

type AddTagYoutrackActionType struct {

}
func(this AddTagYoutrackActionType) Run(
	actionConfig Action, 
	jobContainer JobContainer, 
	clientInterface interface{},
) error {
	var client YoutrackClient
	client = clientInterface.(YoutrackClient)
	arguments := actionConfig.Arguments
	
	youtrackIdVariableName := arguments["youtrackId"]
	tagName := arguments["tagName"]

	youtrackId := jobContainer.Get(youtrackIdVariableName)

	issue, error := client.GetIssue(youtrackId)

	if  error != nil {
		return this.getErrorMessage(error)
	}

	tag := this.createTag(tagName)
	client.AddTagToIssue(&issue, tag)

	error = client.Persist(issue)

	if error != nil {
		return this.getErrorMessage(error)
	}
	return nil
}
func(this AddTagYoutrackActionType) getErrorMessage(err error) error {
	return errors.New(fmt.Sprintf(
		"An error happened in the AddTagYoutrack action for the folowing reason : '%s'", 
		err,
	))
}
func(this AddTagYoutrackActionType) createTag(tagName string) Tag {
	tag := Tag{Name: tagName}
	return tag
}
