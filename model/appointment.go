package model

import (
	"time"
)

type Appointmenter struct {
	ProvinceCity      string    `bson:"Tỉnh/Thành" json:"Tỉnh/Thành"`
	TransactionOffice string    `bson:"PGD" json:"PGD"`
	Service           string    `bson:"Dịch vụ" json:"Dịch vụ"`
	Fullname		  string 	`bson:"Họ tên" json:"Họ tên"`
	PhoneNumber		  int 		`bson:"Điện thoại" json:"Điện thoại"`
	CCCD			  string	`bson:"CCCD" json:"CCCD"`
	Date              time.Time `bson:"Ngày" json:"Ngày"`
	Hour			  time.Time `bson:"Giờ hẹn" json:"Giờ hẹn"`
	Minute			  time.Time  `bson:"Phút hẹn" json:"Phút hẹn"`
	Note 			  string 	`bson:"Ghi chú" json:"Ghi chú"`
}
