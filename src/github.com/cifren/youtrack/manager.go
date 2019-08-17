package youtrack

import (
	"fmt"
)

type ManagerInterface interface {

}

type Manager struct {
	Client Client
}
func (this *Manager) GetIssue(id string) Issue {	
	repo := this.getRepository(Issue{}).(IssueRepository)
	
	return repo.Find(id)
}
func (this *Manager) GetUserTags() []Tag {
	repo := this.getRepository(User{}).(UserRepository)

	return repo.GetMyUser().Tags
}
func (this *Manager) AddTagToUser(tag Tag) {
	repo := this.getRepository(Issue{}).(UserRepository)
	user := repo.GetMyUser()
	user.Tags = append(user.Tags, tag)

	repo.Flush(user)
}
func (this *Manager) AddTagToIssue(issue Issue, tag Tag) {
	issue.Tags = append(issue.Tags, tag)
	repo := this.getRepository(Issue{}).(IssueRepository)

	repo.Flush(issue)
}
func (this *Manager) getRepository(model interface{}) Repository {
	switch model.(type) {
	case User:
		return UserRepository{client: this.Client, route: YoutrackRoutes["my-user"]}
	case Issue:
		return IssueRepository{client: this.Client, route: YoutrackRoutes["issue"]}
	default:
		fmt.Println("That model is not known, failed to find repository")
	}
	return nil
}
