package response

import (
	_ "fmt"
	"go-shared-lib/pkg/meta"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func DefaultErrorResponse(c *fiber.Ctx, err error) error {
	var metaErr meta.MetaError

	if m, ok := meta.IsMetaError(err); ok {
		metaErr = m
	} else {
		m, _ = meta.IsMetaError(meta.ErrorInternalServer.AppendError(err))
		metaErr = m
		metaErr.HttpCode = 500
		metaErr.Message = err.Error()
	}
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())
	// logger.DefaultLogger.Sugar().Errorw(fmt.Sprintf("response method %s path %s status %s", method, path, metaErr.Status),
	// 	"method", method,
	// 	"path", path,
	// 	"error", metaErr.Message,
	// )
	return c.Status(metaErr.GetHTTPCode()).JSON(metaErr)
}

func ResponseBadRequest(c *fiber.Ctx, err error) error {

	var metaErr meta.MetaError

	if m, ok := meta.IsMetaError(err); ok {
		metaErr = m
	} else {

		m, _ = meta.IsMetaError(meta.ErrorBadRequest.AppendError(err))
		metaErr = m
	}
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())

	// logger.DefaultLogger.Sugar().Errorw(fmt.Sprintf("response method %s path %s status %s", method, path, metaErr.Status),
	// 	"method", method,
	// 	"path", path,
	// 	"error", metaErr.Message,
	// )

	return c.Status(metaErr.GetHTTPCode()).JSON(metaErr)
}

func ResponseCreated(c *fiber.Ctx, out interface{}) error {
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())

	// logger.DefaultLogger.Sugar().Infow(fmt.Sprintf("response method %s path %s", method, path),
	// 	"method", method,
	// 	"path", path,
	// 	"response", out,
	// )

	return c.Status(fiber.StatusOK).JSON(out)
}

func ResponseError(c *fiber.Ctx, err error) error {

	var metaErr meta.MetaError

	if m, ok := meta.IsMetaError(err); ok {
		metaErr = m
	} else {
		m, _ = meta.IsMetaError(meta.ErrorInternalServer.AppendError(err))
		metaErr = m
	}
	// method := string(c.Request().Header.Method())
	// 	path := string(c.Request().URI().Path())
	// 	logger.DefaultLogger.Sugar().Errorw(fmt.Sprintf("response method %s path %s status %s", method, path, metaErr.Status),
	// 		"method", method,
	// 		"path", path,
	// 		"error", metaErr.Message,
	// 	)

	return c.Status(metaErr.GetHTTPCode()).JSON(metaErr)
}

func ResponseNoContent(c *fiber.Ctx, out interface{}) error {
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())

	// logger.DefaultLogger.Sugar().Infow(fmt.Sprintf("response method %s path %s", method, path),
	// 	"method", method,
	// 	"path", path,
	// 	"response", out,
	// )

	return c.Status(fiber.StatusNoContent).JSON(out)
}

func ResponseNotFound(c *fiber.Ctx, err error) error {
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())

	var metaErr meta.MetaError

	if m, ok := meta.IsMetaError(err); ok {
		metaErr = m
	} else {
		m, _ = meta.IsMetaError(meta.ErrorNotFound.AppendError(err))
		metaErr = m
	}

	// logger.DefaultLogger.Sugar().Errorw(fmt.Sprintf("response method %s path %s status %s", method, path, metaErr.Status),
	// 	"method", method,
	// 	"path", path,
	// 	"error", metaErr.Message,
	// )

	return c.Status(metaErr.GetHTTPCode()).JSON(metaErr)
}

func ResponseOK(c *fiber.Ctx, out interface{}) error {
	// method := string(c.Request().Header.Method())
	// path := string(c.Request().URI().Path())
	var metaSuccess meta.MetaSuccess
	if reflect.TypeOf(out) == reflect.TypeOf(meta.MetaSuccess{}) {
		metaSuccess = out.(meta.MetaSuccess)
		return c.Status(metaSuccess.HttpCode).JSON(out)
	}
	// logger.DefaultLogger.Sugar().Info(fmt.Sprintf("response method %s path %s", method, path),
	// 	"method", method,
	// 	"path", path,
	// )

	return c.Status(fiber.StatusOK).JSON(out)
}
