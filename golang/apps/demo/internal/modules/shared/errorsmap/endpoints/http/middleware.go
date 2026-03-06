package http

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	domainerrors "github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/domain/errors"
)

type ErrorsMapMiddleware struct {
	logger *slog.Logger
}

func NewErrorsMapMiddleware(logger *slog.Logger) *ErrorsMapMiddleware {
	return &ErrorsMapMiddleware{logger: logger}
}

func (m *ErrorsMapMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		m.logger.ErrorContext(c.Request.Context(), "error",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("error", c.Errors.String()),
		)

		var validationErr validator.ValidationErrors
		if errors.As(err, &validationErr) {
			m.handleValidation(c, &validationErr)

			return
		}

		var domainNotFoundErr *domainerrors.NotFoundError
		if errors.As(err, &domainNotFoundErr) {
			m.handleDomainNotFound(c, domainNotFoundErr)

			return
		}

		var domainForbiddenErr *domainerrors.ForbiddenError
		if errors.As(err, &domainForbiddenErr) {
			m.handleForbidden(c, domainForbiddenErr)

			return
		}

		var domainInvalidArgumentErr *domainerrors.InvalidArgumentError
		if errors.As(err, &domainInvalidArgumentErr) {
			m.handleInvalidArgument(c, domainInvalidArgumentErr)

			return
		}

		var domainPreconditionFailure *domainerrors.PreconditionFailureError
		if errors.As(err, &domainPreconditionFailure) {
			m.handlePreconditionFailure(c, domainPreconditionFailure)

			return
		}

		var domainUnauthenticatedError *domainerrors.UnauthenticatedError
		if errors.As(err, &domainUnauthenticatedError) {
			m.handleUnauthenticated(c, domainUnauthenticatedError)

			return
		}

		m.handleUnknown(c, err)
	}
}

func (m *ErrorsMapMiddleware) handleValidation(c *gin.Context, validationErr *validator.ValidationErrors) {
	type Field struct {
		Rule  string `json:"rule"`
		Field string `json:"field"`
		Param string `json:"param"`
	}

	fields := make([]Field, 0)

	for _, err := range *validationErr {
		f := Field{
			Rule:  err.Tag(),
			Field: err.Namespace(),
			Param: err.Param(),
		}
		fields = append(fields, f)
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fields})
}

func (m *ErrorsMapMiddleware) handleUnknown(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, map[string]any{
		"error": err.Error(),
	})
}

func (m *ErrorsMapMiddleware) handleDomainNotFound(c *gin.Context, err *domainerrors.NotFoundError) {
	c.JSON(http.StatusNotFound, map[string]any{
		"message": err.Error(),
	})
}

func (m *ErrorsMapMiddleware) handleForbidden(c *gin.Context, err *domainerrors.ForbiddenError) {
	c.JSON(http.StatusForbidden, map[string]any{
		"message": err.Error(),
	})
}

func (m *ErrorsMapMiddleware) handleInvalidArgument(c *gin.Context, err *domainerrors.InvalidArgumentError) {
	c.JSON(http.StatusUnprocessableEntity, map[string]any{
		"message": err.Error(),
	})
}

func (m *ErrorsMapMiddleware) handlePreconditionFailure(c *gin.Context, err *domainerrors.PreconditionFailureError) {
	c.JSON(http.StatusPreconditionFailed, map[string]any{
		"message": err.Error(),
	})
}

func (m *ErrorsMapMiddleware) handleUnauthenticated(c *gin.Context, err *domainerrors.UnauthenticatedError) {
	c.JSON(http.StatusUnauthorized, map[string]any{
		"message": err.Error(),
	})
}
