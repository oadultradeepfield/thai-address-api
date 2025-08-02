package queries

import (
	"errors"

	"gorm.io/gorm"
)

type DistrictSortBy uint

const (
	SortByDistrictID DistrictSortBy = iota
	SortByDistrictThaiName
	SortByDistrictEnglishName
)

type DistrictQuery struct {
	BaseQuery

	ProvinceID *uint           `query:"province_id"`
	SortBy     *DistrictSortBy `query:"sort_by"`
}

func (dsb *DistrictSortBy) ToParam() string {
	if dsb == nil || *dsb == SortByDistrictID {
		return "district_id"
	}
	if *dsb == SortByDistrictThaiName {
		return "name_th"
	}
	return "name_en"
}

// Validate checks if the DistrictQuery parameters are valid.
func (dq *DistrictQuery) Validate() error {
	if dq.SortBy != nil &&
		*dq.SortBy != SortByDistrictID &&
		*dq.SortBy != SortByDistrictThaiName &&
		*dq.SortBy != SortByDistrictEnglishName {
		return errors.New("sort_by must be 0 (district_id), 1 (name_th), or 2 (name_en)")
	}
	return dq.BaseQuery.Validate()
}

// Apply applies the DistrictQuery parameters to the given GORM DB instance.
// It is assumed that the district query is validated before calling this method.
func (dq *DistrictQuery) Apply(db *gorm.DB) *gorm.DB {
	db = dq.BaseQuery.Apply(db)
	if dq.ProvinceID != nil {
		db = db.Where("province_id = ?", *dq.ProvinceID)
	}
	if dq.SortBy != nil {
		db = db.Order(dq.SortBy.ToParam() + " " + dq.SortOrder.ToParam())
	}
	return db
}
