package entities

import "time"

type Lead struct {
	ID                     int64
	Name                   string
	Phone                  string
	Source_id              *int64
	Status_id              *int64
	Trial_date             *time.Time
	Qualification          string
	Created_by             *int64
	Created_at             time.Time
	Converted_to_client_at *time.Time
	School_id              string
}

func (l *Lead) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
