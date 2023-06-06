package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	interror "github.com/elraghifary/go-echo-hr-portal/internal/error"
	intsql "github.com/elraghifary/go-echo-hr-portal/internal/sql"
	inttime "github.com/elraghifary/go-echo-hr-portal/internal/time"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := inttime.Init()
	if err != nil {
		interror.ErrorStack(err)
	}
}

func TestEmployeeMySQLRepository(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

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
	repository := employeeMySQLRepository{
		db: db,
		tx: intsql.NewTx(db, identifier.CtxTxKey),
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		wantResult []domain.Employee
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			wantResult: []domain.Employee{
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
			},
			wantErr: false,
			mock: func() {
				columns := []string{
					"id", "nik", "name", "placeOfBirth", "dateOfBirth", "gender", "bloodType", "address", "religion", "maritalStatus", "createdAt", "modifiedAt",
				}
				rows := sqlmock.
					NewRows(columns).
					AddRow(1, "123456789123456", "John Doe", "Jakarta", dateOfBirth, "Laki-laki", "O", "Jalan Baru No. 1", 0, 1, now, nil)
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
			},
			wantResult: []domain.Employee{},
			wantErr:    true,
			mock: func() {
				mock.ExpectQuery("SELECT").WillReturnError(errors.New(identifier.UTAnError))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := repository.Get(test.args.ctx)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.wantResult, got)
		})
	}
}
