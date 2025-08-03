package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"github.com/oadultradeepfield/thai-address-api/internal/repositories"
	"github.com/oadultradeepfield/thai-address-api/internal/responses"
)

func (h *BaseHandler) ListDistrictsHandler(ctx echo.Context) error {
	var query queries.DistrictQuery
	if err := ctx.Bind(&query); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	if err := query.Validate(); err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusBadRequest)
	}

	result, err := repositories.ListDistricts(h.db, &query)
	if err != nil {
		return responses.RespondError(ctx, err.Error(), http.StatusInternalServerError)
	}

	meta := responses.MetaFromQuery(
		&query.BaseQuery,
		result.TotalRecords,
		result.DisplayedRecords,
	)

	response := responses.DistrictResponsesFromModels(result.Districts)
	return responses.RespondSuccess(ctx, meta, response)
}
