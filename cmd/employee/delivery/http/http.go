package http

import (
	"net/http"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	inthttp "github.com/elraghifary/go-echo-hr-portal/internal/response"
	"github.com/labstack/echo/v4"
)

type employeeHTTP struct {
	employeeUsecase domain.EmployeeUsecase
}

func New(e *echo.Echo, employeeUsecase domain.EmployeeUsecase) {
	handler := &employeeHTTP{
		employeeUsecase: employeeUsecase,
	}
	g := e.Group("/api/v1/employee")
	g.GET("", handler.Get)
}

func (h *employeeHTTP) Get(c echo.Context) error {
	ctx := c.Request().Context()

	employees, err := h.employeeUsecase.Get(ctx)
	if err != nil {
		response := inthttp.Response(inthttp.GetStatusCode(err), identifier.ErrGetEmployee.Error(), nil, err.Error())
		return c.JSON(response.Code, response)
	}

	response := inthttp.Response(http.StatusOK, identifier.ResGetEmployee, employees, nil)
	return c.JSON(response.Code, response)
}
