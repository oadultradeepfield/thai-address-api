package queries

import (
	"errors"

	"gorm.io/gorm"
)

type ProvinceSortBy uint

const (
	SortByID ProvinceSortBy = iota
	SortByThaiName
	SortByEnglishName
)

type ProvinceQuery struct {
	BaseQuery

	SortBy *ProvinceSortBy `query:"sort_by"`
}

func (psb *ProvinceSortBy) ToParam() string {
	if psb == nil || *psb == SortByID {
		return "province_id"
	}
	if *psb == SortByThaiName {
		return "name_th"
	}
	return "name_en"
}

// Validate checks if the ProvinceQuery parameters are valid.
func (pq *ProvinceQuery) Validate() error {
	if pq.SortBy != nil &&
		*pq.SortBy != SortByID &&
		*pq.SortBy != SortByThaiName &&
		*pq.SortBy != SortByEnglishName {
		return errors.New("sort_by must be 0 (province_id), 1 (name_th), or 2 (name_en)")
	}
	return pq.BaseQuery.Validate()
}

// Apply applies the ProvinceQuery parameters to the given GORM DB instance.
// It is assumed that the province query is validated before calling this method.
func (pq *ProvinceQuery) Apply(db *gorm.DB) *gorm.DB {
	db = pq.BaseQuery.Apply(db)
	if pq.SortBy != nil {
		db = db.Order(pq.SortBy.ToParam() + " " + pq.SortOrder.ToParam())
	}
	return db
}
