package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/handlers"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
	"gorm.io/gorm"
)

func BaseRoutes(e *echo.Echo, db *gorm.DB) *echo.Echo {
	// Hide server info
	e.HideBanner = true
	e.HidePort = true

	// Register middleware
	Register(e)

	// Health check
	e.GET("/", func(c echo.Context) error {
		return responses.RespondMessage(c, "Service is running")
	})

	// API routes
	api := e.Group("/api/v1")
	h := handlers.NewBaseHandler(db)

	api.GET("/provinces", h.ListProvincesHandler)
	api.GET("/districts", h.ListDistrictsHandler)
	api.GET("/subdistricts", h.ListSubdistrictsHandler)
	api.GET("/subdistricts/:zipcode", h.ListSubdistrictsByZipcodeHandler)

	return e
}
