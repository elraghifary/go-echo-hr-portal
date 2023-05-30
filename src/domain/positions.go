package domain

import "time"

type Position struct {
	Id         int64      `json:"id"`
	UnitID     int64      `json:"unitID"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdAt"`
	ModifiedAt *time.Time `json:"modifiedAt"`
}
