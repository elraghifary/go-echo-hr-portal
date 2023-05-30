package domain

import "time"

type Unit struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdAt"`
	ModifiedAt *time.Time `json:"modifiedAt"`
}
