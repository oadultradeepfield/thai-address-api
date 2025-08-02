package queries

import (
	"errors"

	"gorm.io/gorm"
)

type SortOrder uint

const (
	SortOrderAsc SortOrder = iota
	SortOrderDesc
)

type BaseQuery struct {
	Page      *uint      `query:"page"`
	PageSize  *uint      `query:"page_size"`
	Search    string     `query:"search"`
	SortOrder *SortOrder `query:"sort_order"`
}

func (so *SortOrder) ToParam() string {
	if so == nil || *so == SortOrderAsc {
		return "asc"
	}
	return "desc"
}

// Validate checks if the BaseQuery parameters are valid.
func (bq *BaseQuery) Validate() error {
	if (bq.Page == nil) != (bq.PageSize == nil) {
		return errors.New("page and page_size must both be set or both be unset")
	}
	if bq.Page != nil && *bq.Page < 1 {
		return errors.New("page must be >= 1")
	}
	if bq.PageSize != nil && *bq.PageSize < 1 {
		return errors.New("page_size must be >= 1")
	}
	if bq.SortOrder != nil && *bq.SortOrder != SortOrderAsc && *bq.SortOrder != SortOrderDesc {
		return errors.New("sort_order must be 0 (asc) or 1 (desc)")
	}
	return nil
}

// Apply applies the BaseQuery parameters to the given GORM DB instance.
// It is assumed that the base query is validated before calling this method.
func (bq *BaseQuery) Apply(db *gorm.DB) *gorm.DB {
	if bq.Page != nil && bq.PageSize != nil {
		offset := (*bq.Page - 1) * *bq.PageSize
		db = db.Offset(int(offset)).Limit(int(*bq.PageSize))
	}
	// Search is a common field for filtering by name in provinces, districts, and subdistricts.
	if bq.Search != "" {
		db = db.Where("name_th LIKE ? OR name_en LIKE ?", "%"+bq.Search+"%", "%"+bq.Search+"%")
	}
	// SortOrder is only used together with SortBy in specific queries.
	return db
}
