package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

func ListDistricts(db *gorm.DB, query *queries.DistrictQuery) ([]*models.District, error) {
	var districts []*models.District
	if err := query.Apply(db).Find(&districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}
