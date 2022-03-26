package logger

import "context"

// shared key
const (
	HeaderReferenceID = "x-reference-id"
	HeaderUserID      = "x-user-id"
	HeaderAPIKey      = "x-api-key"
)

func GetFromContext(key string, ctx context.Context) string {
	val, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}

	return val
}
