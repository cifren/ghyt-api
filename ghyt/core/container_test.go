package core

import (
	"testing"
	"github.com/cifren/ghyt-api/ghyt/core/config"
)

func TestRun(t *testing.T) {

	all := make(map[string]interface{})
	all["parameters"] = getParameters()
	all["jobsConfRepository"] = JobRepositoryTest{}

	container := Container{All: all}
	container.InitContainer()
}

type JobRepositoryTest struct {}
func (JobRepositoryTest) GetJobs() ([]config.Job, error) {
  return []config.Job{}, nil
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
