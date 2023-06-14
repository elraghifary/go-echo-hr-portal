package domain

import "time"

type Assignment struct {
	Id         int64
	EmployeeID string
	PositionID string
	StartDate  time.Time
	EndDate    time.Time
	Status     int
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}
