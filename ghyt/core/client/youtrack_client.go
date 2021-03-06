package client

import (
	. "github.com/cifren/ghyt-api/youtrack/core"
	. "github.com/cifren/ghyt-api/youtrack/manager"
	. "github.com/cifren/ghyt-api/youtrack/repository"
)

type YoutrackClient struct {
	manager Manager
	Client Client
}

func NewYoutrackClient(client Client) YoutrackClient {
	youtrackClient := YoutrackClient{}
	youtrackClient.Client = client
	youtrackClient.manager = Manager{Client: youtrackClient.Client}

	return youtrackClient
}

func(this YoutrackClient) GetIssue(id string) (Issue, error) {
	return this.manager.FindIssue(id)
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

	// if doesnt exist in user tags, create tag in youtrack
	tags, err := this.manager.FindTagsByName(tag.Name)
	if err != nil {
	  return err
	}
	if len(tags) == 0 {
		this.manager.Persist("tag", &tag)
	} else {
		tag = tags[0]
	}

	// add tag to issue & flush
	this.manager.AddTagToIssue(issue, tag)
	this.manager.Persist("issue", issue)

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
