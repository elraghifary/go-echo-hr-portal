package domain

import "time"

type Assignment struct {
	Id         int64      `json:"id"`
	EmployeeID string     `json:"employeeID"`
	PositionID string     `json:"positionID"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    time.Time  `json:"endDate"`
	Status     int        `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	ModifiedAt *time.Time `json:"modifiedAt"`
}
