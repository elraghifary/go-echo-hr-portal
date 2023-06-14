package identifier

import "errors"

var (
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrUnauthorized   = errors.New("unauthorized")

	// employee
	ErrGetEmployee = errors.New("failed to get employee")
)
