package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"github.com/oadultradeepfield/thai-address-api/internal/repositories"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
)

func (h *BaseHandler) ListSubdistrictsHandler(ctx echo.Context) error {
	var query queries.SubdistrictQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	if err := query.Validate(); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListSubdistricts(h.db, &query)
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

func (h *BaseHandler) ListSubdistrictsByPostalCodeHandler(ctx echo.Context) error {
	var query queries.PostalCodeQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}
	postalCodeStr := ctx.Param("postal_code")
	if postalCodeUint, err := strconv.ParseUint(postalCodeStr, 10, 0); err == nil {
		postalCode := uint(postalCodeUint)
		query.PostalCode = &postalCode
	} else {
		return responses.RespondError(ctx, "Invalid postal_code parameter", http.StatusBadRequest)
	}

	if err := query.Validate(); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListSubdistrictsByPostalCode(h.db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := &responses.Meta{
		TotalRecords:     result.TotalRecords,
		DisplayedRecords: result.DisplayedRecords,
	}

	response := responses.SubdistrictByPostalCodeResponsesFromModels(result.Subdistricts)
	return responses.RespondSuccess(ctx, meta, response)
}
