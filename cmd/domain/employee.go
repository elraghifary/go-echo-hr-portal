package domain

import (
	"context"
	"time"
)

type (
	EmployeeMySQLRepository interface {
		Get(ctx context.Context) ([]Employee, error)
		Create(ctx context.Context, employee Employee) (int64, error)
		Update(ctx context.Context, employee Employee) (int64, error)
		Delete(ctx context.Context, employee Employee) (int64, error)
	}
)

type Employee struct {
	Id            int64
	NIK           string
	Name          string
	PlaceOfBirth  string
	DateOfBirth   time.Time
	Gender        string
	BloodType     string
	Address       string
	Religion      int
	MaritalStatus int
	CreatedAt     time.Time
	ModifiedAt    *time.Time
}
