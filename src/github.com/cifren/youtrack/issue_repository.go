package youtrack

type IssueRepository struct {
	route string
	client Client
	repository RepositoryHelper
}
func (this *IssueRepository) Find(id string) Issue {
	route := this.route + "/" + id

	request := Request{
		Endpoint: route,
	}
	request.QueryParams.Add("fields", IssueFields)

	var issue Issue
	res, _ := this.client.Get(&request)
	this.repository.Load(res, &issue)
	return issue
}
func (this *IssueRepository) Flush(issue Issue) {
	body := this.repository.GetJson(issue)
	
	request := Request{
		Endpoint: this.route + "/" + issue.Id,
		Body: body,
	}
	request.Headers["Content-Type"] = "application/json"
	
	this.client.Post(&request)
}