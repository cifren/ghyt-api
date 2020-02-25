package manager

import (
	"fmt"
	. "github.com/cifren/ghyt-api/youtrack/repository"
	. "github.com/cifren/ghyt-api/youtrack/core"
)

type Manager struct {
	Client ClientInterface
}

func (this Manager) FindIssue(id string) (Issue, error) {
	repo := this.getRepository("issue").(IssueRepository)
  issue, err := repo.Find(id)

  if err != nil  {
    return Issue{}, err
  }

	return issue.(Issue), nil
}

func (this Manager) FindTagsByName(name string) ([]Tag, error) {
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
		return TagRepository{Client: this.Client, ItemsPerPage: 200}
	case "issue":
		return IssueRepository{Client: this.Client}
	default:
		fmt.Println("That model is not known, failed to find repository")
	}
	return nil
}
