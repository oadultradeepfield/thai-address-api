package models

type District struct {
	ID          uint      `gorm:"column:district_id;primaryKey;not null"`
	ProvinceID  uint      `gorm:"column:province_id"`
	Province    *Province `gorm:"foreignKey:ProvinceID;references:ID"`
	ThaiName    string    `gorm:"column:name_th"`
	EnglishName string    `gorm:"column:name_en"`
}

func (District) TableName() string {
	return "district"
}
