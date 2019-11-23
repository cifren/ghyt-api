package client

import (
	. "github.com/cifren/youtrack"
	// "errors"
)

type YoutrackClient struct {
	Manager Manager
	Client Client
}
func(this YoutrackClient) GetIssue(id string) (Issue, error) {
	issue := this.Manager.GetIssue(id)

	return issue, nil
}
func(this YoutrackClient) Persist(issue Issue) error {
	return nil
}
func(this YoutrackClient) AddTagToIssue(issue *Issue, tag Tag) error {
	name := tag.Name
	
	// check if tag doesn't already exist on the issue
	if _, ok := this.getTag(issue.Tags, name); ok {
		return nil
	}

	// fetch user tags
	userTags := this.Manager.GetUserTags()

	// if doesnt exist in user tags, create tag in youtrack
	if tag, ok := this.getTag(userTags, name); !ok {
		// flush
		user := this.Manager.AddTagToUser(tag)
		this.Manager.Persist(&user)
	}

	// add tag to issue & flush
	this.Manager.AddTagToIssue(issue, tag)
	this.Manager.Persist(issue)

	return nil
}
func (this YoutrackClient) getTag(tags []Tag, name string) (Tag, bool) {
	for key, value := range tags {
		if value.Name == name {
			return tags[key], true
		}
	}
	return Tag{}, false
}