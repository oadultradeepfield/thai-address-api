package responses

import "github.com/oadultradeepfield/thai-address-api/internal/queries"

type Meta struct {
	// Number of records returned in the response
	TotalRecords     uint `json:"total_records"`
	DisplayedRecords uint `json:"displayed_records"`
	// Pagination information (only for paginated responses)
	CurrentPage    uint `json:"current_page,omitempty"`
	RecordsPerPage uint `json:"records_per_page,omitempty"`
	TotalPages     uint `json:"total_pages,omitempty"`
}

func MetaFromQuery(
	query *queries.BaseQuery,
	totalRecords uint,
	displayedRecords uint,
) Meta {
	meta := Meta{
		TotalRecords:     totalRecords,
		DisplayedRecords: displayedRecords,
	}
	if query.Page != nil && query.PageSize != nil {
		meta.CurrentPage = *query.Page
		meta.RecordsPerPage = *query.PageSize
		meta.TotalPages = (totalRecords + *query.Page - 1) / *query.PageSize
	}
	return meta
}
