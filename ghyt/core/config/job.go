package config

type Job struct {
	Conditions []Condition `json:"conditions"`
	Actions []Action `json:"actions"`
	Id int `json:"id"`
}
