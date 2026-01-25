package main

import (
	appErrors "study-stack/internal/shared/app_errors"

	"github.com/gofiber/fiber/v2"
)

// maps ServiceError to generic HTTP status messages
func FiberErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).SendString(e.Message)
	}

	var code int
	var msg string
	if se, ok := err.(appErrors.ServiceError); ok {
		switch se {
		case appErrors.InternalError, appErrors.NotImplemented, appErrors.ServiceUnavailable:
			code = fiber.StatusInternalServerError
			msg = "Internal Server Error"
		case appErrors.BadData, appErrors.ValidationFailed, appErrors.InvalidCredentials:
			code = fiber.StatusBadRequest
			msg = "Bad Request"
		case appErrors.Unauthorized:
			code = fiber.StatusUnauthorized
			msg = "Unauthorized"
		case appErrors.Forbidden:
			code = fiber.StatusForbidden
			msg = "Forbidden"
		case appErrors.NotFound:
			code = fiber.StatusNotFound
			msg = "Not Found"
		case appErrors.MethodNotAllowed:
			code = fiber.StatusMethodNotAllowed
			msg = "Method Not Allowed"
		case appErrors.Conflict:
			code = fiber.StatusConflict
			msg = "Conflict"
		case appErrors.TooManyRequests:
			code = fiber.StatusTooManyRequests
			msg = "Too Many Requests"
		}

	}

	return c.Status(code).SendString(msg)
}
