package config

type JobRepositoryInterface interface {
  GetJobs() []Job
}
