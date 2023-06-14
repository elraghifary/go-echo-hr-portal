package response

import (
	"net/http"

	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
)

type ResponseJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func Response(code int, message string, data, errors interface{}) ResponseJSON {
	res := ResponseJSON{}
	res.Code = code
	res.Message = message
	res.Data = data
	res.Errors = errors

	return res
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.Error() {
	case identifier.ErrUnauthorized.Error():
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
