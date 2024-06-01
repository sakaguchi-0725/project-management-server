package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/pkg/apperr"
	"go.uber.org/zap"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func makeCode(err apperr.AppError) string {
	return fmt.Sprintf("%s/%s", apperr.MakeType(err), apperr.MakeCategory(err))
}

func ErrorMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			logger.Error("API Request Error",
				zap.String("Error", err.Error()),
				zap.String("Method", c.Request().Method),
				zap.String("URL", c.Request().URL.Path),
			)

			if appErr, ok := err.(apperr.AppError); ok {
				switch appErr.Type() {
				case apperr.ErrForbidden:
					return c.NoContent(http.StatusForbidden)
				case apperr.ErrUnauthorized:
					return c.JSON(http.StatusUnauthorized, &ErrorResponse{
						Code: makeCode(appErr),
					})
				case apperr.ErrBadRequest:
					return c.JSON(http.StatusBadRequest, &ErrorResponse{
						Code:    makeCode(appErr),
						Message: appErr.Error(),
					})
				case apperr.ErrNotFound:
					return c.JSON(http.StatusNotFound, &ErrorResponse{
						Code:    makeCode(appErr),
						Message: appErr.Error(),
					})
				}
			}

			return c.NoContent(http.StatusInternalServerError)
		}
	}
}
