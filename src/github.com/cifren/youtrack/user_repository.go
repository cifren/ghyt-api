package youtrack

type UserRepository struct {
	route string
	client Client
	repository RepositoryHelper
}
func (this UserRepository) GetMyUser() User {
	return this.Find("").(User)
}
// Id is empty at anytime
func (this UserRepository) Find(id string) interface{} {
	return this.FindUser(id)
}
func (this UserRepository) FindUser(id string) User {
	route := this.route

	request := Request{
		Endpoint: route,
	}
	request.QueryParams.Add("fields", UserFields)

	res, _ := this.client.Get(&request)
	var user User
	
	this.repository.Load(res, &user)

	return user
}
func (this UserRepository) Flush(user interface{}) {
	myUser := user.(User)
	this.FlushUser(&myUser)
}
func (this UserRepository) FlushUser(user *User) {
	body := this.repository.GetJson(user)

	request := Request{
		Endpoint: this.route,
		Body: body,
	}
	request.Headers["Content-Type"] = "application/json"

	this.client.Post(&request)
}