package middleware

import (
	"time"
	"go-backend-task/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func SetupMiddleware(app *fiber.App) {
	// 1. Request ID Middleware
	app.Use(requestid.New())

	// 2. Custom Zap Logger Middleware
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		
		// Process request
		err := c.Next()

		// Log details
		logger.Log.Info("Incoming Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("request_id", c.GetRespHeader("X-Request-ID")),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
		)

		return err
	})
}