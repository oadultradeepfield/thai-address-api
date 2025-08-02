package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

func ListProvinces(db *gorm.DB, query *queries.ProvinceQuery) ([]*models.Province, error) {
	var provinces []*models.Province
	if err := query.Apply(db).Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
