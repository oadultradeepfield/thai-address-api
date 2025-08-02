package models

type Subdistrict struct {
	ID          uint      `gorm:"column:subdistrict_id;primaryKey"`
	DistrictID  uint      `gorm:"column:district_id"`
	District    *District `gorm:"foreignKey:DistrictID;references:ID"`
	ThaiName    string    `gorm:"column:name_th"`
	EnglishName string    `gorm:"column:name_en"`
	PostalCode  uint      `gorm:"column:postal_code"`
}

func (Subdistrict) TableName() string {
	return "subdistrict"
}
