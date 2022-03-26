package errmsg

import "go-shared-lib/pkg/meta"

var (
	ErrInvalid_Page_Format  = meta.ErrorBadRequest.SetHTTPCode(400).SetCode(1001).AppendMessage("page is invalid format, It should be number.")
	ErrInvalid_Limit_Format = meta.ErrorBadRequest.SetHTTPCode(400).SetCode(1002).AppendMessage("limit is invalid format, It should be number.")
	ErrInvalid_Sort_Format  = meta.ErrorBadRequest.SetHTTPCode(400).SetCode(1003).AppendMessage("sort is invalid format, It should be [{field}-asc,{field}-desc].")
	ErrInvalidContentType   = meta.ErrorBadRequest.SetCode(422).AppendMessage("Invalid content-type")
	ErrInternal             = meta.ErrorInternalServer.SetCode(500).AppendMessage("Internal Error")
	ErrAdminNotAllow        = meta.ErrorForbidden.AppendMessage("require admin-id in header")
	ErrAdminDenied          = meta.ErrorForbidden.AppendMessage("permission denied")
)

// DB error
var (
	ErrForeignKey     = meta.ErrorBadRequest.SetCode(1452).AppendMessage("ER_NO_REFERENCED_ROW: Cannot add or update a child row,a foreign key constraint fails.")
	ErrColumnInvalid  = meta.ErrorBadRequest.SetCode(1054).AppendMessage("ER_BAD_FIELD_ERROR")
	ErrDeniedAccess   = meta.ErrorInternalServer.SetCode(1142).AppendMessage("ER_TABLEACCESS_DENIED_ERROR")
	ErrDatabase       = meta.ErrorInternalServer.AppendMessage("DATABASE_ERROR")
	ErrUpdateNotFound = meta.ErrorBadRequest.SetCode(40010).AppendMessage("not found record to update")
	ErrDeleteNotFound = meta.ErrorBadRequest.SetCode(40020).AppendMessage("not found record to delete")
)
