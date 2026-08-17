package dto

type UpdateWardRequest struct {
	Name       *string `json:"name" form:"name"`
	Code       *string `json:"code" form:"code"`
	ProvinceID *string `json:"province_id" form:"province_id"`
}

type UpdateWardDTO = UpdateWardRequest
