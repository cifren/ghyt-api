package event

type GithubEvent interface {
	SetName(name string)
	GetName() string
	SetVariables([]VariableInterface)
	GetVariables() []VariableInterface
}

type VariableInterface interface {
	SetName(name string) VariableInterface
	GetName() string
	SetValue(val string) VariableInterface
	GetValue() string
}

type PullRequestEvent struct {
	Variables map[string]string
}
func NewPullRequestEvent() *PullRequestEvent {
	return &PullRequestEvent{Variables: make(map[string]string)}
}
func (this *PullRequestEvent) SetVariables(variables map[string]string) {
	this.Variables = variables
}
func (this *PullRequestEvent) GetVariables() map[string]string {
	return this.Variables
}