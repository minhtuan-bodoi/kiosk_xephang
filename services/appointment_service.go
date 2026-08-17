package services

import (
	"errors"
	"fmt"
	"strings"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateAppointment(appointment model.Appointmenter) error {
	if strings.TrimSpace(appointment.CCCD) == "" {
		return fmt.Errorf("%w: CCCD không được để trống", errors.New("Dữ liệu không hợp lệ"))
	}
	if strings.TrimSpace(appointment.Fullname) == "" {
		return fmt.Errorf("%w: họ tên không được để trống", errors.New("Dữ liệu không hợp lệ"))
	}

	return repositories.CreateAppointment(appointment)
}

func GetAllAppointments() ([]model.Appointmenter, error) {
	return repositories.GetAppointments()
}

func GetAppointmentByCCCD(cccd string) (model.Appointmenter, error) {
	return repositories.GetAppointmentByCCCD(cccd)
}

func UpdateAppointment(id bson.ObjectID, req dto.UpdateAppointmentRequest) error {
	updates := bson.M{}

	if req.ProvinceCity != nil {
		val := strings.TrimSpace(*req.ProvinceCity)
		if val == "" {
			return fmt.Errorf("%w: tỉnh/thành không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["Tỉnh/Thành"] = val
	}
	if req.TransactionOffice != nil {
		val := strings.TrimSpace(*req.TransactionOffice)
		if val == "" {
			return fmt.Errorf("%w: PGD không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["PGD"] = val
	}
	if req.Service != nil {
		val := strings.TrimSpace(*req.Service)
		if val == "" {
			return fmt.Errorf("%w: dịch vụ không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["Dịch vụ"] = val
	}
	if req.Fullname != nil {
		val := strings.TrimSpace(*req.Fullname)
		if val == "" {
			return fmt.Errorf("%w: họ tên không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["Họ tên"] = val
	}
	if req.PhoneNumber != nil {
		updates["Điện thoại"] = *req.PhoneNumber
	}
	if req.CCCD != nil {
		val := strings.TrimSpace(*req.CCCD)
		if val == "" {
			return fmt.Errorf("%w: CCCD không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["CCCD"] = val
	}
	if req.Date != nil {
		updates["Ngày"] = *req.Date
	}
	if req.Hour != nil {
		updates["Giờ hẹn"] = *req.Hour
	}
	if req.Minute != nil {
		updates["Phút hẹn"] = *req.Minute
	}
	if req.Note != nil {
		updates["Ghi chú"] = strings.TrimSpace(*req.Note)
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	return repositories.UpdateAppointment(id, updates)
}

func DeleteAppointment(id string) error {
	return repositories.DeleteAppointment(id)
}
