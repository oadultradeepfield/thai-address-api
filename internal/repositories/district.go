package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

// DistrictListResult encapsulates the result of listing districts
type DistrictListResult struct {
	Districts        []*models.District
	TotalRecords     uint
	DisplayedRecords uint
}

func ListDistricts(db *gorm.DB, query *queries.DistrictQuery) (*DistrictListResult, error) {
	// Get total count before applying filters/pagination
	var totalRecords int64
	if err := db.Model(&models.District{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	// Apply query filters and retrieve districts
	var districts []*models.District
	queryDB := query.Apply(db)
	if err := queryDB.Find(&districts).Error; err != nil {
		return nil, err
	}

	return &DistrictListResult{
		Districts:        districts,
		TotalRecords:     uint(totalRecords),
		DisplayedRecords: uint(len(districts)),
	}, nil
}
