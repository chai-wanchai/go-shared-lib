package meta

import (
	"net/http"
)

type MetaSuccess struct {
	HttpCode int         `json:"-"`
	Message  string      `json:"-"`
	Status   bool        `json:"status" example:"true"`
	Data     interface{} `json:"data,omitempty" swaggerignore:"true"`
	Meta     interface{} `json:"meta,omitempty" swaggerignore:"true"`
}

var (
	Success        = newMetaSuccess(http.StatusOK)
	CreatedSuccess = newMetaSuccess(http.StatusCreated)
)

func newMetaSuccess(httpCode int) MetaSuccess {
	return MetaSuccess{
		HttpCode: httpCode,
		Status:   true,
	}
}
func (m MetaSuccess) SetHTTPCode(code int) MetaSuccess {
	m.HttpCode = code
	return m
}

func (m MetaSuccess) SetStatus(status bool) MetaSuccess {
	m.Status = status
	return m
}

func (m MetaSuccess) SetMessage(msg string) MetaSuccess {
	m.Message = msg
	return m
}

func (m MetaSuccess) SetData(data interface{}) MetaSuccess {
	m.Data = data
	return m
}
func (m MetaSuccess) SetMeta(meta interface{}) MetaSuccess {
	m.Meta = meta
	return m
}
