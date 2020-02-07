package core

import (
	"testing"
	"github.com/cifren/ghyt-api/ghyt/core/config"
)

func TestRun(t *testing.T) {

	all := make(map[string]interface{})
	all["parameters"] = getParameters()
	all["jobRepository"] = JobRepositoryTest{}

	container := Container{All: all}
	container.InitContainer()
}

type JobRepositoryTest struct {}
func (JobRepositoryTest) GetJobs() []config.Job {
  return []config.Job {}
}

func getParameters() map[string]interface{} {

	return map[string]interface{}{
        "github": map[string]string{
            "github_account": "github_account-test",
            "github_secret": "github_secret-test",
        },
        "youtrack": map[string]string{
            "url": "url-test",
            "token": "token-test",
        },
    }
}
