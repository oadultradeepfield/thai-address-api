package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

// ProvinceListResult encapsulates the result of listing provinces
type ProvinceListResult struct {
	Provinces        []*models.Province
	TotalRecords     uint
	DisplayedRecords uint
}

func ListProvinces(db *gorm.DB, query *queries.ProvinceQuery) (*ProvinceListResult, error) {
	// Get total count before applying filters/pagination
	var totalRecords int64
	if err := db.Model(&models.Province{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	// Apply query filters and retrieve provinces
	var provinces []*models.Province
	queryDB := query.Apply(db)
	if err := queryDB.Find(&provinces).Error; err != nil {
		return nil, err
	}

	return &ProvinceListResult{
		Provinces:        provinces,
		TotalRecords:     uint(totalRecords),
		DisplayedRecords: uint(len(provinces)),
	}, nil
}
