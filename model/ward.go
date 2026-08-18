package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Ward struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Code         string        `bson:"code" json:"code"`
	ProvinceCode string        `bson:"province_id" json:"province_id"`
	ProvinceName string        `bson:"province_name" json:"province_name"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
}
