package dto

import "mime/multipart"

type UpdateServiceRequest struct {
	NameServiceVN *string               `form:"name_service_VN" json:"name_service_VN"`
	NameServiceEN *string               `form:"name_service_EN" json:"name_service_EN"`
	CodeService   *string               `form:"code_service" json:"code_service"`
	IconService   *multipart.FileHeader `form:"icon_service" json:"-"`
}

type UpdateServiceDTO = UpdateServiceRequest


