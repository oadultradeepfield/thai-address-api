package models

type Subdistrict struct {
	ID          uint      `gorm:"column:subdistrict_id;primaryKey"`
	DistrictID  uint      `gorm:"column:district_id"`
	District    *District `gorm:"foreignKey:DistrictID;references:ID"`
	ThaiName    string    `gorm:"column:name_th"`
	EnglishName string    `gorm:"column:name_en"`
	Zipcode     uint      `gorm:"column:zipcode"`
}

func (Subdistrict) TableName() string {
	return "subdistrict"
}
