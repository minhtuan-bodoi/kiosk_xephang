package dto

type UpdateTicketRequest struct {
	TicketCode       *string `form:"ticket_code" json:"ticket_code"`
	TicketFormatCode *string `form:"ticket_format_code" json:"ticket_format_code"`
}

type UpdateTicketDTO = UpdateTicketRequest

