package repository

import (
	. "github.com/cifren/youtrack/core"
)

const (
	ISSUE_ENDPOINT string = "issues"
)

type IssueRepository struct {
	Client Client
}

func (this IssueRepository) Find(id string) interface {} {
	return this.FindIssue(id)
}

func (this IssueRepository) FindIssue(id string) Issue {
	endpoint := ISSUE_ENDPOINT + "/" + id
	issue := Issue{}

	this.getRepository().Find(&issue, endpoint, this.Client, IssueFields)

	return issue
}

func (this IssueRepository) Flush(issuePointer interface{}) {
	myIssue := issuePointer.(*Issue)
	this.FlushIssue(myIssue)
}

func (this IssueRepository) FlushIssue(issue *Issue) {
	endpoint := ISSUE_ENDPOINT + "/" + (*issue).Id

	this.getRepository().Flush(&issue, endpoint, this.Client)
}

func(this IssueRepository) getRepository() RepositoryHelper {
    return RepositoryHelper{}
}
