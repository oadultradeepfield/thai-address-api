package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"github.com/oadultradeepfield/thai-address-api/internal/repositories"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
	"gorm.io/gorm"
)

func ListDistrictsHandler(ctx echo.Context, db *gorm.DB) error {
	var query queries.DistrictQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListDistricts(db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := &responses.Meta{
		TotalRecords:     result.TotalRecords,
		DisplayedRecords: result.DisplayedRecords,
	}

	response := responses.DistrictResponsesFromModels(result.Districts)
	return responses.RespondSuccess(ctx, meta, response)
}
