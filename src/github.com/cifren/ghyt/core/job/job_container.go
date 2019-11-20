package job

type JobContainer struct {
	Variables map[string]string
}
func NewJobContainer() *JobContainer {
	v := new(JobContainer)	
	v.Variables = make(map[string]string)
	return v
}
func(this JobContainer) Get(name string) string {
	return this.Variables[name]
}
func(this *JobContainer) Set(name string, value string) {
	this.Variables[name] = value
}