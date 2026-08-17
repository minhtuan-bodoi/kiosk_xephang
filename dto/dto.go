package dto

import (
	"mime/multipart"
)

type UpdateTicketRequest struct {
	TicketCode       *string `form:"ticket_code" json:"ticket_code"`
	TicketFormatCode *string `form:"ticket_format_code" json:"ticket_format_code"`
}

type UpdateServiceRequest struct {
	NameServiceVN *string               `form:"name_service_VN" json:"name_service_VN"`
	NameServiceEN *string               `form:"name_service_EN" json:"name_service_EN"`
	CodeService   *string               `form:"code_service" json:"code_service"`
	IconService   *multipart.FileHeader `form:"icon_service" json:"-"`
}
