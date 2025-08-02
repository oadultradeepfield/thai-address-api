package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register sets up all global middlewares for the Echo instance
func Register(e *echo.Echo) {
	e.Use(
		middleware.Logger(),
		middleware.Recover(),

		// Basic security headers
		middleware.Secure(),

		// CORS
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods: []string{echo.GET},
			AllowOrigins: []string{"*"},
		}),

		// Rate limiting (60 requests per minute per IP)
		middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(60)),

		// Request timeout
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 10 * time.Second,
		}),

		// Request size limit
		middleware.BodyLimit("1MB"),

		// Compression
		middleware.Gzip(),
	)
}
