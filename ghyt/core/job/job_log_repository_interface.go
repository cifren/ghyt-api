package job

type JobLogRepositoryInterface interface {
  Persist(log JobLog) (error)
}
