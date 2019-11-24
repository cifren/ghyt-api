package manager

import (
	"fmt"
	. "github.com/cifren/youtrack/repository"
	. "github.com/cifren/youtrack/core"
)

type Manager struct {
	Client Client
}

func (this Manager) FindIssue(id string) Issue {	
	var repo IssueRepository = this.getRepository("issue").(IssueRepository)
	
	return repo.Find(id).(Issue)
}

func (this Manager) FindTagsByName(name string) []Tag {
	repo := this.getRepository("tag").(TagRepository)

	return repo.FindTagsByName(name)
}

func (this Manager) AddTagToIssue(issue *Issue, tag Tag) {
	issue.Tags = append(issue.Tags, tag)
}

func (this Manager) Persist(modelName string, modelPointer interface{}) {
	repository := this.getRepository(modelName)
	repository.Flush(modelPointer)
}

func (this Manager) getRepository(modelName string) RepositoryInterface {
	switch modelName {
	case "tag":
		return TagRepository{client: this.Client}
	case "issue":
		return IssueRepository{client: this.Client}
	default:
		fmt.Println("That model is not known, failed to find repository")
	}
	return nil
}
