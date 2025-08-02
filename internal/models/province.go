package models

type Province struct {
	ID          uint   `gorm:"column:province_id;primaryKey;not null"`
	ThaiName    string `gorm:"column:name_th;not null"`
	EnglishName string `gorm:"column:name_en;not null"`
}

func (Province) TableName() string {
	return "province"
}
