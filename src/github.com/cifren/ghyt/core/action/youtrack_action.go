package action

import (
	"github.com/cifren/youtrack"
)

type ActionInterface interface {
	GetName() string
	Run(client youtrack.Client)
}

type AddTag struct {
	Name string
}
func (this AddTag) GetName() string {
	return this.Name
}
func (this AddTag) Run(client youtrack.Client) {
	manager := youtrack.Manager{Client: client}
	
	// get issue tags
	issue := manager.GetIssue("connect-1517")
	
	// get user tags
	userTags := manager.GetUserTags()
	
	// check if tag doesn't already exist
	if _, ok := this.getTag(issue.Tags, this.Name); ok {
		return
	}

	tag, ok := this.getTag(userTags, this.Name); 
	// if doesnt exist in user tags, create tag
	if !ok {
		// flush
		manager.AddTagToUser(tag)
	}
	
	// add tag to issue & flush
	manager.AddTagToIssue(issue, tag)
}
func (this *AddTag) getTag(tags []youtrack.Tag, name string) (youtrack.Tag, bool) {
	for key, value := range tags {
		if value.Name == name {
			return tags[key], true
		}
	}
	return youtrack.Tag{}, false
}

type ActionRunner struct {

}
func (this *ActionRunner) RunIt(action ActionInterface) {
	youtrackUrl := "https://cospirit.myjetbrains.com/youtrack/api"
	token := "Bearer perm:RnJhbmNpc19MZWNvcQ==.R2h5dA==.d9qeAMnrZ2V8JGs7DGICy6aijn3wWg"
	client := youtrack.Client{Url: youtrackUrl, Token: token}
	
	action.Run(client)
}
