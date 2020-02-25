package config

type Condition struct {
	Name string `json:"name"`
	Arguments map[string]string `json:"arguments"`
	Id int `json:"id"`
}
