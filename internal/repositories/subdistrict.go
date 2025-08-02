package repositories

import (
	"github.com/oadultradeepfield/thai-address-api/internal/models"
	"github.com/oadultradeepfield/thai-address-api/internal/queries"
	"gorm.io/gorm"
)

func ListSubdistricts(db *gorm.DB, query *queries.SubdistrictQuery) ([]*models.Subdistrict, error) {
	var subdistricts []*models.Subdistrict
	if err := query.Apply(db).Find(&subdistricts).Error; err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func ListSubdistrictsByZipcode(db *gorm.DB, query *queries.ZipcodeQuery) ([]*models.Subdistrict, error) {
	var subdistricts []*models.Subdistrict
	if err := PreloadAssociations(query.Apply(db)).
		Find(&subdistricts).Error; err != nil {
		return nil, err
	}
	return subdistricts, nil
}

func PreloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("District").
		Preload("District.Province")
}
