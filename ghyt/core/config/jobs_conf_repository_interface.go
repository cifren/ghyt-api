package config

type JobsConfRepositoryInterface interface {
  GetJobs() ([]Job, error)
}
