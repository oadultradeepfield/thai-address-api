package queries

import (
	"errors"

	"gorm.io/gorm"
)

type SubdistrictSortBy uint

const (
	SortBySubdistrictID SubdistrictSortBy = iota
	SortBySubdistrictThaiName
	SortBySubdistrictEnglishName
	SortBySubdistrictPostalCode
)

type SubdistrictQuery struct {
	BaseQuery

	DistrictID *uint              `query:"district_id"`
	SortBy     *SubdistrictSortBy `query:"sort_by"`
}

type PostalCodeQuery struct {
	BaseQuery

	PostalCode *uint              `query:"postal_code"`
	SortBy     *SubdistrictSortBy `query:"sort_by"`
}

func (ssb *SubdistrictSortBy) ToParam() string {
	if ssb == nil || *ssb == SortBySubdistrictID {
		return "subdistrict_id"
	}
	if *ssb == SortBySubdistrictThaiName {
		return "name_th"
	}
	if *ssb == SortBySubdistrictEnglishName {
		return "name_en"
	}
	return "postal_code"
}

// Validate checks if the SubdistrictQuery parameters are valid.
func (sq *SubdistrictQuery) Validate() error {
	if sq.SortBy != nil &&
		*sq.SortBy != SortBySubdistrictID &&
		*sq.SortBy != SortBySubdistrictThaiName &&
		*sq.SortBy != SortBySubdistrictEnglishName &&
		*sq.SortBy != SortBySubdistrictPostalCode {
		return errors.New("sort_by must be 0 (subdistrict_id), 1 (name_th), 2 (name_en), or 3 (postal_code)")
	}
	return sq.BaseQuery.Validate()
}

// Validate checks if the PostalCodeQuery parameters are valid.
func (pcq *PostalCodeQuery) Validate() error {
	return pcq.BaseQuery.Validate()
}

// Apply applies the SubdistrictQuery parameters to the given GORM DB instance.
// It is assumed that the subdistrict query is validated before calling this method.
func (sq *SubdistrictQuery) Apply(db *gorm.DB) *gorm.DB {
	db = sq.BaseQuery.Apply(db)
	if sq.DistrictID != nil {
		db = db.Where("district_id = ?", *sq.DistrictID)
	}
	if sq.SortBy != nil {
		db = db.Order(sq.SortBy.ToParam() + " " + sq.SortOrder.ToParam())
	}
	return db
}

// Apply applies the PostalCodeQuery parameters to the given GORM DB instance.
// It is assumed that the postal code query is validated before calling this method.
func (pcq *PostalCodeQuery) Apply(db *gorm.DB) *gorm.DB {
	db = pcq.BaseQuery.Apply(db)
	if pcq.PostalCode != nil {
		db = db.Where("postal_code = ?", *pcq.PostalCode)
	}
	if pcq.SortBy != nil {
		db = db.Order(pcq.SortBy.ToParam() + " " + pcq.SortOrder.ToParam())
	}
	return db
}
