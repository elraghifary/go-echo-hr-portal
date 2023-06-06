package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	mockdomain "github.com/elraghifary/go-echo-hr-portal/cmd/mock"
	interror "github.com/elraghifary/go-echo-hr-portal/internal/error"
	inttime "github.com/elraghifary/go-echo-hr-portal/internal/time"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := inttime.Init()
	if err != nil {
		interror.ErrorStack(err)
	}
}

func TestEmployeeUsecase(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	now := time.Date(2023, 5, 24, 0, 0, 0, 0, inttime.GetLocation())
	tempNow := inttime.Now
	inttime.Now = func() time.Time {
		return now
	}
	defer func() {
		inttime.Now = tempNow
	}()
	dateOfBirth := time.Date(2006, 1, 2, 0, 0, 0, 0, inttime.GetLocation())
	employeeMySQLRepository := mockdomain.NewMockEmployeeMySQLRepository(ctrl)
	uc := &employeeUsecase{
		employeeMySQLRepository: employeeMySQLRepository,
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		wantResult []domain.EmployeeResponse
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			wantResult: []domain.EmployeeResponse{
				{
					Id:            1,
					NIK:           "123456789123456",
					Name:          "John Doe",
					PlaceOfBirth:  "Jakarta",
					DateOfBirth:   dateOfBirth,
					Gender:        "Laki-laki",
					BloodType:     "O",
					Address:       "Jalan Baru No. 1",
					Religion:      0,
					MaritalStatus: 1,
				},
			},
			wantErr: false,
			mock: func() {
				employeeMySQLRepository.EXPECT().Get(gomock.Any()).Return([]domain.Employee{
					{
						Id:            1,
						NIK:           "123456789123456",
						Name:          "John Doe",
						PlaceOfBirth:  "Jakarta",
						DateOfBirth:   dateOfBirth,
						Gender:        "Laki-laki",
						BloodType:     "O",
						Address:       "Jalan Baru No. 1",
						Religion:      0,
						MaritalStatus: 1,
						CreatedAt:     now,
						ModifiedAt:    nil,
					},
				}, nil)
			},
		},
		{
			name: "success empty",
			args: args{
				ctx: ctx,
			},
			wantResult: []domain.EmployeeResponse{},
			wantErr:    false,
			mock: func() {
				employeeMySQLRepository.EXPECT().Get(gomock.Any()).Return([]domain.Employee{}, nil)
			},
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				employeeMySQLRepository.EXPECT().Get(gomock.Any()).Return([]domain.Employee{}, errors.New(identifier.UTAnError))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := uc.Get(test.args.ctx)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.wantResult, got)
		})
	}
}
