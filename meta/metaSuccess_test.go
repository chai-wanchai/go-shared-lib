package meta

import (
	"reflect"
	"testing"
)

func TestMetaSuccess_SetHTTPCode(t *testing.T) {
	type fields struct {
		HttpCode int
		Message  string
		Status   bool
		Data     interface{}
		Meta     interface{}
	}
	type args struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaSuccess
	}{
		{
			name:   "MetaSuccess.SetHTTPCode()",
			fields: fields{},
			args:   args{code: 9000},
			want:   MetaSuccess{HttpCode: 9000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaSuccess{
				HttpCode: tt.fields.HttpCode,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				Data:     tt.fields.Data,
				Meta:     tt.fields.Meta,
			}
			if got := m.SetHTTPCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaSuccess.SetHTTPCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaSuccess_SetStatus(t *testing.T) {
	type fields struct {
		HttpCode int
		Message  string
		Status   bool
		Data     interface{}
		Meta     interface{}
	}
	type args struct {
		status bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaSuccess
	}{
		{
			name:   "MetaSuccess.SetStatus()",
			fields: fields{},
			args:   args{status: false},
			want:   MetaSuccess{Status: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaSuccess{
				HttpCode: tt.fields.HttpCode,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				Data:     tt.fields.Data,
				Meta:     tt.fields.Meta,
			}
			if got := m.SetStatus(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaSuccess.SetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaSuccess_SetMessage(t *testing.T) {
	type fields struct {
		HttpCode int
		Message  string
		Status   bool
		Data     interface{}
		Meta     interface{}
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaSuccess
	}{
		{
			name:   "MetaSuccess.SetMessage()",
			fields: fields{},
			args:   args{msg: "kk"},
			want:   MetaSuccess{Message: "kk"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaSuccess{
				HttpCode: tt.fields.HttpCode,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				Data:     tt.fields.Data,
				Meta:     tt.fields.Meta,
			}
			if got := m.SetMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaSuccess.SetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaSuccess_SetData(t *testing.T) {
	type fields struct {
		HttpCode int
		Message  string
		Status   bool
		Data     interface{}
		Meta     interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaSuccess
	}{
		{
			name:   "MetaSuccess.SetData()",
			fields: fields{},
			args:   args{data: "kk"},
			want:   MetaSuccess{Data: "kk"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaSuccess{
				HttpCode: tt.fields.HttpCode,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				Data:     tt.fields.Data,
				Meta:     tt.fields.Meta,
			}
			if got := m.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaSuccess.SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaSuccess_SetMeta(t *testing.T) {
	type fields struct {
		HttpCode int
		Message  string
		Status   bool
		Data     interface{}
		Meta     interface{}
	}
	type args struct {
		meta interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MetaSuccess
	}{
		{
			name:   "MetaSuccess.SetData()",
			fields: fields{},
			args:   args{meta: "kk"},
			want:   MetaSuccess{Meta: "kk"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetaSuccess{
				HttpCode: tt.fields.HttpCode,
				Message:  tt.fields.Message,
				Status:   tt.fields.Status,
				Data:     tt.fields.Data,
				Meta:     tt.fields.Meta,
			}
			if got := m.SetMeta(tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetaSuccess.SetMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
