package responses

import "github.com/oadultradeepfield/thai-address-api/internal/models"

type SubdistrictResponse struct {
	SubdistrictID uint   `json:"subdistrict_id"`
	ThaiName      string `json:"thai_name"`
	EnglishName   string `json:"english_name"`
	PostalCode    uint   `json:"postal_code"`
}

// A response for a subdistrict, including its district and province.
// This is used when looking up a subdistrict by its postal code.
type SubdistrictByPostalCodeResponse struct {
	SubdistrictResponse

	District *DistrictResponse `json:"district"`
	Province *ProvinceResponse `json:"province"`
}

func SubdistrictResponseFromModel(model *models.Subdistrict) *SubdistrictResponse {
	return &SubdistrictResponse{
		SubdistrictID: model.ID,
		ThaiName:      model.ThaiName,
		EnglishName:   model.EnglishName,
		PostalCode:    model.PostalCode,
	}
}

func SubdistrictResponsesFromModels(models []*models.Subdistrict) []*SubdistrictResponse {
	responses := make([]*SubdistrictResponse, len(models))
	for i, model := range models {
		responses[i] = SubdistrictResponseFromModel(model)
	}
	return responses
}

// It is assumed that model preloads the District and Province relations.
func SubdistrictByPostalCodeResponseFromModel(model *models.Subdistrict) *SubdistrictByPostalCodeResponse {
	return &SubdistrictByPostalCodeResponse{
		SubdistrictResponse: *SubdistrictResponseFromModel(model),
		District:            DistrictResponseFromModel(model.District),
		Province:            ProvinceResponseFromModel(model.District.Province),
	}
}

func SubdistrictByPostalCodeResponsesFromModels(models []*models.Subdistrict) []*SubdistrictByPostalCodeResponse {
	responses := make([]*SubdistrictByPostalCodeResponse, len(models))
	for i, model := range models {
		responses[i] = SubdistrictByPostalCodeResponseFromModel(model)
	}
	return responses
}
