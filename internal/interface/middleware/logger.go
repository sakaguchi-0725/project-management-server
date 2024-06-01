package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			// リクエストボディ読み取り
			reqBody, _ := io.ReadAll(req.Body)
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			// 次のハンドラを呼び出す
			err := next(c)

			status := res.Status
			length := res.Size

			logger.Info("API Request",
				zap.String("Method", req.Method),
				zap.String("Path", req.URL.Path),
				zap.String("IP", c.RealIP()),
				zap.ByteString("Request Body", reqBody),
				zap.Int("Status", status),
				zap.Int64("Length", length),
				zap.Duration("Duration", time.Since(start)),
			)

			return err
		}
	}
}
