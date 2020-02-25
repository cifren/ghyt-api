package job

type JobLog struct {
  Request map[string]interface{} `json:"request"`
  Feedback []JobFeedback `json:"feedback"`
  ErrorMessage string `json:"error_message"`
}
