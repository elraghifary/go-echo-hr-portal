package main

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	employeemysql "github.com/elraghifary/go-echo-hr-portal/cmd/employee/repository/mysql"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	interror "github.com/elraghifary/go-echo-hr-portal/internal/error"
	intsql "github.com/elraghifary/go-echo-hr-portal/internal/sql"
	inttime "github.com/elraghifary/go-echo-hr-portal/internal/time"
	"github.com/labstack/echo/v4"
)

var (
	mysql                   *sql.DB
	tx                      *intsql.Tx
	employeeMySQLRepository domain.EmployeeMySQLRepository
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
	employeeMySQLRepository = employeemysql.New(mysql, tx)
}

func initHTTP() error {
	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return err
	}

	e := echo.New()
	e.Debug = debug

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

	err = initHTTP()
	if err != nil {
		interror.ErrorStack(err)
	}
}
