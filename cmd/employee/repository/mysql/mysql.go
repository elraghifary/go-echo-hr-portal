package mysql

import (
	"context"
	"database/sql"

	"github.com/elraghifary/go-echo-hr-portal/cmd/domain"
	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
	intsql "github.com/elraghifary/go-echo-hr-portal/internal/sql"
	"github.com/sirupsen/logrus"
)

type employeeMySQLRepository struct {
	db *sql.DB
	tx *intsql.Tx
}

func New(db *sql.DB, tx *intsql.Tx) domain.EmployeeMySQLRepository {
	return &employeeMySQLRepository{
		db: db,
		tx: tx,
	}
}

func (r *employeeMySQLRepository) fetchAll(ctx context.Context, query string, args ...interface{}) ([]domain.Employee, error) {
	employees := []domain.Employee{}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		employee := domain.Employee{}
		err = rows.Scan(
			&employee.Id,
			&employee.NIK,
			&employee.Name,
			&employee.PlaceOfBirth,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.BloodType,
			&employee.Address,
			&employee.Religion,
			&employee.MaritalStatus,
			&employee.CreatedAt,
			&employee.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	return employees, nil
}

func (r *employeeMySQLRepository) fetchOne(ctx context.Context, query string, args ...interface{}) (*domain.Employee, error) {
	employee := domain.Employee{}
	row := r.db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&employee.Id,
		&employee.NIK,
		&employee.Name,
		&employee.PlaceOfBirth,
		&employee.DateOfBirth,
		&employee.Gender,
		&employee.BloodType,
		&employee.Address,
		&employee.Religion,
		&employee.MaritalStatus,
		&employee.CreatedAt,
		&employee.ModifiedAt,
	)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeMySQLRepository) execContext(ctx context.Context, query string, args ...interface{}) (identifier.LastInsertId, identifier.RowsAffected, error) {
	var db interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	} = r.db

	tx, _ := ctx.Value(r.tx.GetTxKey()).(*sql.Tx)
	if tx != nil {
		db = tx
	}

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	if rowsAffected == 0 {
		return 0, 0, identifier.ErrNoRowsAffected
	}

	return lastInsertId, rowsAffected, nil
}

func (r *employeeMySQLRepository) Get(ctx context.Context) ([]domain.Employee, error) {
	result, err := r.fetchAll(ctx, QueryGet)
	if err != nil {
		return []domain.Employee{}, err
	}

	return result, nil
}

func (r *employeeMySQLRepository) Create(ctx context.Context, employee domain.Employee) (int64, error) {
	id, _, err := r.execContext(ctx, QueryCreate,
		employee.CreatedAt,
		employee.ModifiedAt,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *employeeMySQLRepository) Update(ctx context.Context, employee domain.Employee) (int64, error) {
	_, rows, err := r.execContext(ctx, QueryUpdate,
		employee.ModifiedAt,
		employee.Id,
	)
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (r *employeeMySQLRepository) Delete(ctx context.Context, employee domain.Employee) (int64, error) {
	_, rows, err := r.execContext(ctx, QueryDelete,
		employee.Id,
	)
	if err != nil {
		return 0, err
	}

	return rows, nil
}
