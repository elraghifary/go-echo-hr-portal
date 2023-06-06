package usecase

import (
	"context"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	intsql "github.com/elraghifary/go-echo-hr-portal/internal/sql"
)

type employeeUsecase struct {
	employeeMySQLRepository domain.EmployeeMySQLRepository
	tx                      *intsql.Tx
}

func New(employeeMySQLRepository domain.EmployeeMySQLRepository, tx *intsql.Tx) domain.EmployeeUsecase {
	return &employeeUsecase{
		employeeMySQLRepository: employeeMySQLRepository,
		tx:                      tx,
	}
}

func (u *employeeUsecase) Get(ctx context.Context) ([]domain.EmployeeResponse, error) {
	employees, err := u.employeeMySQLRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	employeeResponse := []domain.EmployeeResponse{}
	for i := range employees {
		employee := domain.EmployeeResponse{
			Id:            employees[i].Id,
			NIK:           employees[i].NIK,
			Name:          employees[i].Name,
			PlaceOfBirth:  employees[i].PlaceOfBirth,
			DateOfBirth:   employees[i].DateOfBirth,
			Gender:        employees[i].Gender,
			BloodType:     employees[i].BloodType,
			Address:       employees[i].Address,
			Religion:      employees[i].Religion,
			MaritalStatus: employees[i].MaritalStatus,
		}

		employeeResponse = append(employeeResponse, employee)
	}

	return employeeResponse, nil
}
