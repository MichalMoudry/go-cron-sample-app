package ioc

import "gocron-sample/database/model"

// An interface for a job result's repository.
type IJobResultRepository interface {
	// Method for obtaining a collection of job results in the database. Default row limit is 1000.
	GetJobResults(limit int) (model.JobResult, error)
}
