package main

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	employeehttpdelivery "github.com/elraghifary/go-echo-hr-portal/cmd/employee/delivery/http"
	employeemysqlrepository "github.com/elraghifary/go-echo-hr-portal/cmd/employee/repository/mysql"
	employeeusecase "github.com/elraghifary/go-echo-hr-portal/cmd/employee/usecase"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	interror "github.com/elraghifary/go-echo-hr-portal/internal/error"
	intsql "github.com/elraghifary/go-echo-hr-portal/internal/sql"
	inttime "github.com/elraghifary/go-echo-hr-portal/internal/time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mysql                   *sql.DB
	tx                      *intsql.Tx
	employeeMySQLRepository domain.EmployeeMySQLRepository
	employeeUsecase         domain.EmployeeUsecase
)

func initOther() error {
	err := inttime.Init()
	tx = intsql.NewTx(mysql, identifier.CtxTxKey)

	return err
}

func initDB() error {
	mysqlConn, err := intsql.OpenMySQL(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	mysql = mysqlConn
	return err
}

func initRepository() {
	employeeMySQLRepository = employeemysqlrepository.New(mysql, tx)
}

func initUsecase() {
	employeeUsecase = employeeusecase.New(employeeMySQLRepository, tx)
}

func initHTTPDelivery() error {
	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return err
	}

	e := echo.New()
	e.Debug = debug
	e.Use(middleware.CORS())

	employeehttpdelivery.New(e, employeeUsecase)

	return e.Start(os.Getenv("SERVER_PORT"))
}

func main() {
	err := initOther()
	if err != nil {
		interror.ErrorStack(err)
	}

	err = initDB()
	if err != nil {
		interror.ErrorStack(err)
	}
	defer func() {
		if mysql != nil {
			err = mysql.Close()
			if err != nil {
				interror.ErrorStack(err)
			}
		}
	}()

	initRepository()
	initUsecase()

	err = initHTTPDelivery()
	if err != nil {
		interror.ErrorStack(err)
	}
}
