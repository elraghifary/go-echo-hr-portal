package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	mockdomain "github.com/elraghifary/go-echo-hr-portal/cmd/mock"
	interror "github.com/elraghifary/go-echo-hr-portal/internal/error"
	inttime "github.com/elraghifary/go-echo-hr-portal/internal/time"
)

func init() {
	err := inttime.Init()
	if err != nil {
		interror.ErrorStack(err)
	}
}

func TestNew(t *testing.T) {
	New(echo.New(), nil)
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Date(2023, 5, 24, 0, 0, 0, 0, inttime.GetLocation())
	tempNow := inttime.Now
	inttime.Now = func() time.Time {
		return now
	}
	defer func() {
		inttime.Now = tempNow
	}()
	dateOfBirth := time.Date(2006, 1, 2, 0, 0, 0, 0, inttime.GetLocation())
	employeeMock := mockdomain.NewMockEmployeeUsecase(ctrl)
	deliveryHTTP := &employeeHTTP{
		employeeUsecase: employeeMock,
	}

	tests := []struct {
		name       string
		args       string
		wantResult string
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: "",
			wantResult: `{"code":200,"message":"` + identifier.ResGetEmployee + `","data":[{"id":1,"nik":"123456789123456","name":"John Doe","placeOfBirth":"Jakarta","dateOfBirth":"02 Jan 2006","gender":"Laki-laki","bloodType":"O","address":"Jalan Baru No. 1","religion":"Islam","maritalStatus":"Belum Kawin"}],"errors":null}
`,
			wantErr: false,
			mock: func() {
				employeeResponse := []domain.EmployeeResponse{
					{
						Id:            1,
						NIK:           "123456789123456",
						Name:          "John Doe",
						PlaceOfBirth:  "Jakarta",
						DateOfBirth:   dateOfBirth.Format(identifier.TimeFormat5),
						Gender:        "Laki-laki",
						BloodType:     "O",
						Address:       "Jalan Baru No. 1",
						Religion:      "Islam",
						MaritalStatus: "Belum Kawin",
					},
				}
				employeeMock.EXPECT().Get(gomock.Any()).Return(employeeResponse, nil)
			},
		},
		{
			name: "error",
			args: "",
			wantResult: `{"code":500,"message":"` + identifier.ErrGetEmployee.Error() + `","data":null,"errors":"` + identifier.UTAnError + `"}
`,
			wantErr: true,
			mock: func() {
				employeeMock.EXPECT().Get(gomock.Any()).Return(nil, errors.New(identifier.UTAnError))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			req := httptest.NewRequest(http.MethodGet, identifier.UTPathEmployee, nil)
			req.Header.Set(identifier.UTContentType, identifier.UTContentTypeJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			if assert.NoError(t, deliveryHTTP.Get(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}
