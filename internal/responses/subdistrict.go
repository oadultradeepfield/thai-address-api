package responses

import "github.com/oadultradeepfield/thai-address-api/internal/models"

type SubdistrictResponse struct {
	SubdistrictID uint   `json:"subdistrict_id"`
	ThaiName      string `json:"thai_name"`
	EnglishName   string `json:"english_name"`
	Zipcode       uint   `json:"zipcode"`
}

// A response for a subdistrict, including its district and province.
// This is used when looking up a subdistrict by its zipcode.
type SubdistrictByZipcodeResponse struct {
	SubdistrictResponse

	District *DistrictResponse `json:"district"`
	Province *ProvinceResponse `json:"province"`
}

func SubdistrictResponseFromModel(model *models.Subdistrict) *SubdistrictResponse {
	return &SubdistrictResponse{
		SubdistrictID: model.ID,
		ThaiName:      model.ThaiName,
		EnglishName:   model.EnglishName,
		Zipcode:       model.Zipcode,
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
func SubdistrictByZipcodeResponseFromModel(model *models.Subdistrict) *SubdistrictByZipcodeResponse {
	return &SubdistrictByZipcodeResponse{
		SubdistrictResponse: *SubdistrictResponseFromModel(model),
		District:            DistrictResponseFromModel(model.District),
		Province:            ProvinceResponseFromModel(model.District.Province),
	}
}

func SubdistrictByZipcodeResponsesFromModels(models []*models.Subdistrict) []*SubdistrictByZipcodeResponse {
	responses := make([]*SubdistrictByZipcodeResponse, len(models))
	for i, model := range models {
		responses[i] = SubdistrictByZipcodeResponseFromModel(model)
	}
	return responses
}
