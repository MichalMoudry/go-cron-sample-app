package model

import (
	"time"

	"github.com/google/uuid"
)

type JobResult struct {
	Id       uuid.UUID
	Name     string
	Status   string
	Started  time.Time
	Finished time.Time
}
