package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

// SubdistrictListResult encapsulates the result of listing subdistricts
type SubdistrictListResult struct {
	Subdistricts     []*models.Subdistrict
	TotalRecords     uint
	DisplayedRecords uint
}

// ListSubdistricts retrieves subdistricts based on the provided query and returns a result struct.
func ListSubdistricts(db *gorm.DB, query *queries.SubdistrictQuery) (*SubdistrictListResult, error) {
	// Get total count before applying filters/pagination
	var totalRecords int64
	if err := db.Model(&models.Subdistrict{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	// Apply query filters and retrieve subdistricts
	var subdistricts []*models.Subdistrict
	queryDB := query.Apply(db)
	if err := queryDB.Find(&subdistricts).Error; err != nil {
		return nil, err
	}

	return &SubdistrictListResult{
		Subdistricts:     subdistricts,
		TotalRecords:     uint(totalRecords),
		DisplayedRecords: uint(len(subdistricts)),
	}, nil
}

// ListSubdistrictsByPostalCode retrieves subdistricts by postal code and returns a result struct.
func ListSubdistrictsByPostalCode(db *gorm.DB, query *queries.PostalCodeQuery) (*SubdistrictListResult, error) {
	// Get total count before applying filters/pagination
	var totalRecords int64
	if err := db.Model(&models.Subdistrict{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	// Apply query filters and retrieve subdistricts
	var subdistricts []*models.Subdistrict
	queryDB := query.Apply(db)
	preloadDB := PreloadAssociations(queryDB)
	if err := preloadDB.Find(&subdistricts).Error; err != nil {
		return nil, err
	}

	return &SubdistrictListResult{
		Subdistricts:     subdistricts,
		TotalRecords:     uint(totalRecords),
		DisplayedRecords: uint(len(subdistricts)),
	}, nil
}

func PreloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("District").
		Preload("District.Province")
}
