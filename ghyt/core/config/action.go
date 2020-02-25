package config

type Action struct {
	To string `json:"to"`
	Name string `json:"name"`
	Arguments map[string]string `json:"arguments"`
	Id int `json:"id"`
}
