package domain

import "time"

type Position struct {
	Id        int64
	UnitID    int64
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
