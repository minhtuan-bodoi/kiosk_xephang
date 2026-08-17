package dto

import "time"

type UpdateAppointmentRequest struct {
	ProvinceCity      *string    `form:"Tỉnh/Thành" json:"Tỉnh/Thành"`
	TransactionOffice *string    `form:"PGD" json:"PGD"`
	Service           *string    `form:"Dịch vụ" json:"Dịch vụ"`
	Fullname          *string    `form:"Họ tên" json:"Họ tên"`
	PhoneNumber       *int       `form:"Điện thoại" json:"Điện thoại"`
	CCCD              *string    `form:"CCCD" json:"CCCD"`
	Date              *time.Time `form:"Ngày" json:"Ngày"`
	Hour              *time.Time `form:"Giờ hẹn" json:"Giờ hẹn"`
	Minute            *time.Time `form:"Phút hẹn" json:"Phút hẹn"`
	Note              *string    `form:"Ghi chú" json:"Ghi chú"`
}
