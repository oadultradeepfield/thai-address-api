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
	SortBySubdistrictZipcode
)

type SubdistrictQuery struct {
	BaseQuery

	DistrictID *uint              `query:"district_id"`
	SortBy     *SubdistrictSortBy `query:"sort_by"`
}

type ZipcodeQuery struct {
	BaseQuery

	Zipcode *uint              `query:"zipcode"`
	SortBy  *SubdistrictSortBy `query:"sort_by"`
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
	return "zipcode"
}

// Validate checks if the SubdistrictQuery parameters are valid.
func (sq *SubdistrictQuery) Validate() error {
	if sq.SortBy != nil &&
		*sq.SortBy != SortBySubdistrictID &&
		*sq.SortBy != SortBySubdistrictThaiName &&
		*sq.SortBy != SortBySubdistrictEnglishName &&
		*sq.SortBy != SortBySubdistrictZipcode {
		return errors.New("sort_by must be 0 (subdistrict_id), 1 (name_th), 2 (name_en), or 3 (zipcode)")
	}
	return sq.BaseQuery.Validate()
}

// Validate checks if the ZipcodeQuery parameters are valid.
func (zq *ZipcodeQuery) Validate() error {
	return zq.BaseQuery.Validate()
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

// Apply applies the ZipcodeQuery parameters to the given GORM DB instance.
// It is assumed that the zipcode query is validated before calling this method.
func (zq *ZipcodeQuery) Apply(db *gorm.DB) *gorm.DB {
	db = zq.BaseQuery.Apply(db)
	if zq.Zipcode != nil {
		db = db.Where("zipcode = ?", *zq.Zipcode)
	}
	if zq.SortBy != nil {
		db = db.Order(zq.SortBy.ToParam() + " " + zq.SortOrder.ToParam())
	}
	return db
}
