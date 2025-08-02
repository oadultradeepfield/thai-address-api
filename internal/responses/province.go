package responses

import "github.com/oadultradeepfield/thai-address-api/internal/models"

type ProvinceResponse struct {
	ProvinceID  uint   `json:"province_id"`
	ThaiName    string `json:"thai_name"`
	EnglishName string `json:"english_name"`
}

func ProvinceResponseFromModel(model *models.Province) *ProvinceResponse {
	return &ProvinceResponse{
		ProvinceID:  model.ID,
		ThaiName:    model.ThaiName,
		EnglishName: model.EnglishName,
	}
}
