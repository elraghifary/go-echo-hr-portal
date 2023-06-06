// Package error implements some error utility functions
package error

import (
	errorsst "github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

// Returns stack traces of error
func ErrorStack(err error) {
	man := errorsst.New(err)
	logrus.Error(man.ErrorStack())
	// logrus.Error(err.(*errorsst.Error).ErrorStack())
}
