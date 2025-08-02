package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"github.com/oadultradeepfield/thai-address-api/internal/repositories"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
	"gorm.io/gorm"
)

func ListProvincesHandler(ctx echo.Context, db *gorm.DB) error {
	var query queries.ProvinceQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListProvinces(db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := &responses.Meta{
		TotalRecords:     result.TotalRecords,
		DisplayedRecords: result.DisplayedRecords,
	}

	response := responses.ProvinceResponsesFromModels(result.Provinces)
	return responses.RespondSuccess(ctx, meta, response)
}
