package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	NameService      string `bson:"name_service" json:"name_service"`
	CodeService      string `bson:"code_service" json:"code_service" `
	TicketFormatCode string `bson:"ticket_format_code" json:"ticket_format_code"`
	IconService      string `bson:"icon_service" json:"icon_service"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
