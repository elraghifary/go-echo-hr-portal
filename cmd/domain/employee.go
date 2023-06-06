package domain

import (
	"context"
	"time"
)

type (
	EmployeeMySQLRepository interface {
		Get(ctx context.Context) ([]Employee, error)
	}

	EmployeeUsecase interface {
		Get(ctx context.Context) ([]EmployeeResponse, error)
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

type EmployeeResponse struct {
	Id            int64     `json:"id"`
	NIK           string    `json:"nik"`
	Name          string    `json:"name"`
	PlaceOfBirth  string    `json:"placeOfBirth"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	Gender        string    `json:"gender"`
	BloodType     string    `json:"bloodType"`
	Address       string    `json:"address"`
	Religion      int       `json:"religion"`
	MaritalStatus int       `json:"maritalStatus"`
}
