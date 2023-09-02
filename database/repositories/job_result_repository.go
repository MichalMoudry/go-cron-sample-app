package repositories

import (
	"gocron-sample/database"
	"gocron-sample/database/model"
	"gocron-sample/database/queries"
)

type JobResultRepository struct{}

// Method for obtaining a collection of job results in the database.
func GetJobResults(limit int) ([]model.JobResult, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return make([]model.JobResult, 0), err
	}
	if limit != 0 {
		var jobResults []model.JobResult
		if err = ctx.Get(&jobResults, queries.GetJobResults, limit); err != nil {
			return make([]model.JobResult, 0), err
		}
		return jobResults, nil
	}
	// TODO: Add unlimited query.
	return make([]model.JobResult, 0), nil
}
