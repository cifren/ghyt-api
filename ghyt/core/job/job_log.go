package job

import (
  "time"
)

type JobLog struct {
  CreatedAt time.Time `json:"created_at"`
  Request map[string]interface{} `json:"request"`
  Feedback []JobFeedback `json:"feedback"`
  ErrorMessage string `json:"error_message"`
}
