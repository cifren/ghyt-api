package client

import (
	. "github.com/cifren/youtrack"
	// "errors"
)

type YoutrackClient struct {

}
func(this YoutrackClient) GetIssue(id string) (Issue, error) {
	return Issue{}, nil
}
func(this YoutrackClient) Persist(issue Issue) error {
	return nil
}
func(this YoutrackClient) AddTagToIssue(issue *Issue, tag Tag) error {
	return nil
}