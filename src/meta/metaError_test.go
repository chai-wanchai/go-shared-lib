package meta

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newMetaError(t *testing.T) {
	type args struct {
		httpCode int
	}
	tests := []struct {
		name string
		args args
		want MetaError
	}{
		{name: "Init Success",
			args: args{httpCode: 500},
			want: MetaError{HttpCode: 500, Code: -1000, Message: "", Status: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMetaError(tt.args.httpCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMetaError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaError_SetHTTPCode(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaError
	}{
		{
			name: "HTTP 400",
			fields: fields{
				HttpCode: 400,
			},
			args: args{code: 400},
			want: MetaError{
				HttpCode: 400,
				Message:  "",
				Status:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			if got := m.SetHTTPCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaError.SetHTTPCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaError_SetStatus(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	type args struct {
		status bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaError
	}{
		{
			name: "SetStatus to true",
			fields: fields{
				HttpCode: 500,
			},
			args: args{status: true},
			want: MetaError{
				HttpCode: 500,
				Status:   true,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			if got := m.SetStatus(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaError.SetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaError_SetCode(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaError
	}{
		{
			name: "SetCode",
			fields: fields{
				HttpCode: 500,
			},
			args: args{code: 9000},
			want: MetaError{
				HttpCode: 500,
				Code:     9000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			if got := m.SetCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaError.SetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaError_GetHTTPCode(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Get Http code",
			fields: fields{HttpCode: 500},
			want:   500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			if got := m.GetHTTPCode(); got != tt.want {
				t.Errorf("MetaError.GetHTTPCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaError_AppendMessage(t *testing.T) {
	assert := assert.New(t)
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaError
	}{
		{
			name:   "Add Error Message",
			fields: fields{},
			args:   args{format: "1_%v", a: []interface{}{"2"}},
			want: MetaError{
				Message: "1_2",
			},
		},
		{
			name: "Append Error Message",
			fields: fields{
				Message: "Err :",
			},
			args: args{format: "1"},
			want: MetaError{
				Message: "Err :1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			got := m.AppendMessage(tt.args.format, tt.args.a...)
			assert.Equal(tt.want.Message, got.Message)

		})
	}
}

func TestMetaError_AppendError(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	type args struct {
		err error
	}
	assert := assert.New(t)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaError
	}{
		{
			name:   "Add Errors",
			fields: fields{err: errors.New("Internal")},
			args:   args{err: errors.New("Internal")},
			want: MetaError{
				Message: "Internal",
				err:     errors.New("Internal"),
			},
		},
		{
			name:   "Add Errors , Check Error",
			fields: fields{Code: 0},
			args:   args{err: MetaError{Code: 7000, Message: "unit"}},
			want: MetaError{
				Message: "unit",
				Code:    7000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			got := m.AppendError(tt.args.err)
			assert.Equal(tt.want.Message, got.Message)
			assert.Equal(tt.want.Code, got.Code)

		})
	}
}

func TestMetaError_Error(t *testing.T) {
	type fields struct {
		HttpCode int
		Code     int
		Message  string
		Status   bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Error()",
			fields: fields{
				HttpCode: 3000,
				Code:     3000,
			},
			want: `{"http_code":3000,"code":3000,"message":"","status":false}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaError{
				HttpCode: tt.fields.HttpCode,
				Code:     tt.fields.Code,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				err:      tt.fields.err,
			}
			if got := m.Error(); got != tt.want {
				t.Errorf("MetaError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
