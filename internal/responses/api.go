package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	// Success respond will have Meta and Data fields
	Meta *Meta `json:"meta,omitempty"`
	Data any   `json:"data,omitempty"`
	// Message and error respond will have Message field
	Message string `json:"message,omitempty"`
}

func RespondSuccess(
	ctx echo.Context,
	meta *Meta,
	data any,
) error {
	response := &ApiResponse{
		Meta: meta,
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

func RespondMessage(
	ctx echo.Context,
	message string,
) error {
	response := &ApiResponse{
		Message: message,
	}
	return ctx.JSON(http.StatusOK, response)
}

func RespondError(
	ctx echo.Context,
	errorMessage string,
	status int,
) error {
	response := &ApiResponse{
		Message: errorMessage,
	}
	return ctx.JSON(status, response)
}
