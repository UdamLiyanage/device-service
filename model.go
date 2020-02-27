package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Serial       string             `json:"serial" bson:"serial"`
	UID          string             `json:"uid" bson:"uid"`
	DefID        string             `json:"device_definition_id" bson:"device_definition_id"`
	Firmware     string             `json:"firmware_version" bson:"firmware_version"`
	LastFirmware time.Time          `json:"last_firmware_update" bson:"last_firmware_update"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
