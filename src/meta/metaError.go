package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// DefaultCode is the default error code value if not defined
const DefaultCode = -1000

type MetaError struct {
	HttpCode int    `json:"http_code" example:"500"`
	Code     int    `json:"code" example:"-1000"`
	Message  string `json:"message" example:"Message Error"`
	Status   bool   `json:"status" example:"false"`

	err error
}

var (
	// client error responses
	ErrorBadRequest         = newMetaError(http.StatusBadRequest)
	ErrorUnauthorized       = newMetaError(http.StatusUnauthorized)
	ErrorPreconditionFailed = newMetaError(http.StatusPreconditionFailed)
	ErrorForbidden          = newMetaError(http.StatusForbidden)
	ErrorNotFound           = newMetaError(http.StatusNotFound)
	ErrorRequestTimeout     = newMetaError(http.StatusRequestTimeout)

	// server error responses
	ErrorInternalServer = newMetaError(http.StatusInternalServerError)
	ErrorNotImplemented = newMetaError(http.StatusNotImplemented)
)

func IsMetaError(err error) (metaErr MetaError, ok bool) {
	ok = errors.As(err, &metaErr)
	return metaErr, ok
}

func newMetaError(httpCode int) MetaError {
	return MetaError{
		HttpCode: httpCode,
		Code:     DefaultCode,
		Status:   false,
	}
}
func (m MetaError) GetHTTPCode() int {
	return m.HttpCode
}

func (m MetaError) SetHTTPCode(code int) MetaError {
	m.HttpCode = code
	return m
}
func (m MetaError) SetCode(code int) MetaError {
	m.Code = code
	return m
}
func (m MetaError) SetStatus(status bool) MetaError {
	m.Status = status
	return m
}

func (m MetaError) AppendMessage(format string, a ...interface{}) MetaError {
	m.err = fmt.Errorf(format, a...)
	if m.Message != "" {
		m.Message = fmt.Sprint(m.Message, format)
	} else {
		m.Message = fmt.Sprintf(format, a...)
	}

	return m
}

func (m MetaError) AppendError(err error) MetaError {

	if metaErr, ok := IsMetaError(err); ok {
		if m.Code == 0 {
			m.Code = metaErr.Code
		}

		err = errors.New(metaErr.Message)
	}

	m.err = err
	m.Message = err.Error()

	return m
}
func (m MetaError) Error() string {
	b, _ := json.Marshal(m)
	return string(b)
}
