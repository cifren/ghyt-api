package repository

import (
	. "github.com/cifren/ghyt-api/youtrack/core"
)

const (
	ISSUE_ENDPOINT string = "issues"
)

type IssueRepository struct {
	Client ClientInterface
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
	endpoint := ISSUE_ENDPOINT + "/" + (*issue).IdReadable
	this.getRepository().Flush(issue, endpoint, this.Client, IssueFields, nil)
}

func(this IssueRepository) getRepository() RepositoryHelper {
    return RepositoryHelper{}
}
