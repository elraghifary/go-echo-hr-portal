package domain

import "time"

type Unit struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
