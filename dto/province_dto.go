package dto

type UpdateProvinceRequest struct {
	Name *string `json:"name" form:"name"`
	Code *string `json:"code" form:"code"`
}

type UpdateProvinceDTO = UpdateProvinceRequest
