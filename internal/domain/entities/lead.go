package entities

import "time"

type Lead struct {
	ID                  int64
	Name                string
	Phone               string
	SourceID            int64
	StatusID            int64
	TrialDate           time.Time
	Qualification       string
	CreatedBy           int64
	CreatedAt           time.Time
	ConvertedToClientAt time.Time
	SchoolID            string
}
