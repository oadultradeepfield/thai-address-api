package responses

import "github.com/oadultradeepfield/thai-address-api/internal/models"

type DistrictResponse struct {
	DistrictID  uint   `json:"district_id"`
	ThaiName    string `json:"thai_name"`
	EnglishName string `json:"english_name"`
}

func DistrictResponseFromModel(model *models.District) *DistrictResponse {
	return &DistrictResponse{
		DistrictID:  model.ID,
		ThaiName:    model.ThaiName,
		EnglishName: model.EnglishName,
	}
}
