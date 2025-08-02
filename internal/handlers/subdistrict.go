package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"github.com/oadultradeepfield/thai-address-api/internal/repositories"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
	"gorm.io/gorm"
)

func ListSubdistrictsHandler(ctx echo.Context, db *gorm.DB) error {
	var query queries.SubdistrictQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListSubdistricts(db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := &responses.Meta{
		TotalRecords:     result.TotalRecords,
		DisplayedRecords: result.DisplayedRecords,
	}

	response := responses.SubdistrictResponsesFromModels(result.Subdistricts)
	return responses.RespondSuccess(ctx, meta, response)
}

func ListSubdistrictsByZipcodeHandler(ctx echo.Context, db *gorm.DB) error {
	var query queries.ZipcodeQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListSubdistrictsByZipcode(db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := &responses.Meta{
		TotalRecords:     result.TotalRecords,
		DisplayedRecords: result.DisplayedRecords,
	}

	response := responses.SubdistrictByZipcodeResponsesFromModels(result.Subdistricts)
	return responses.RespondSuccess(ctx, meta, response)
}
