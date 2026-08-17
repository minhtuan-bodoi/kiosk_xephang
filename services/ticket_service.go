package services

import (
	"errors"
	"fmt"
	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/repositories"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTicket(ticket model.Ticket) error {

	if strings.TrimSpace(ticket.TicketCode) == "" {
		return fmt.Errorf("%w: mã phòng ban không được để trống", errors.New("Dữ liệu không hợp lệ"))
	}
	if strings.TrimSpace(ticket.TicketFormatCode) == "" {
		return fmt.Errorf("%w: mã phòng ban không được để trống", errors.New("Dữ liệu không hợp lệ"))
	}

	exists, err := repositories.IsServiceCodeExists(ticket.TicketFormatCode)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%w: mã phòng ban đã tồn tại", errors.New("Dữ liệu không hợp lệ"))
	}

	return repositories.CreateTicket(ticket)
}

func GetAllTickets() ([]model.Ticket, error) {
	return repositories.GetTickets()
}

func GetTicketByID(id string) (model.Ticket, error) {
	return repositories.GetTicketByID(id)
}

func GetTicketByTicketCode(code string) (model.Ticket, error) {
	return repositories.GetTicketByTicketCode(code)
}

func UpdateTicket(id bson.ObjectID, req dto.UpdateTicketRequest) error {
	updates := bson.M{}

	if req.TicketCode != nil {
		code := strings.ToUpper(strings.TrimSpace(*req.TicketCode))
		if code == "" {
			return fmt.Errorf("%w: mã không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["ticket_code"] = code
	}
	if req.TicketFormatCode != nil {
		codeFormat := strings.TrimSpace(*req.TicketFormatCode)
		if codeFormat == "" {
			return fmt.Errorf("%w: mã định dạng vé không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["ticket_format_code"] = codeFormat
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	return repositories.UpdateTicket(id, updates)
}

func DeleteTicket(id string) error {
	return repositories.DeleteTicket(id)
}
