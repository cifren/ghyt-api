package repository

import (
	. "github.com/cifren/youtrack/core"
)

const (
	ISSUE_ENDPOINT string = "issues"
)

type IssueRepository struct {
	client Client
	repository RepositoryHelper
}
func (this IssueRepository) Find(id string) interface {} {
	return this.FindIssue(id)
}
func (this IssueRepository) FindIssue(id string) Issue {
	endpoint := ISSUE_ENDPOINT + "/" + id
	issue := Issue{}
	
	this.repository.Find(&issue, endpoint, this.client, IssueFields)

	return issue
}
func (this IssueRepository) Flush(issuePointer interface{}) {
	myIssue := issuePointer.(*Issue)
	this.FlushIssue(myIssue)
}
func (this IssueRepository) FlushIssue(issue *Issue) {
	endpoint := ISSUE_ENDPOINT + "/" + (*issue).Id

	this.repository.Flush(&issue, endpoint, this.client)
}