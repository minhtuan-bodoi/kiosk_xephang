package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Ticket struct {
	ID               bson.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketCode       string        `bson:"ticket_code" json:"ticket_code" `
	TicketFormatCode string        `bson:"ticket_format_code" json:"ticket_format_code" `
	CreatedAt        time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time     `bson:"updated_at" json:"updated_at"`
}
