package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Serial       string             `json:"serial" bson:"serial"`
	UUID         string             `json:"uuid" bson:"uuid"`
	DefID        string             `json:"device_definition_id" bson:"device_definition_id"`
	Firmware     string             `json:"firmware_version" bson:"firmware_version"`
	LastFirmware primitive.DateTime `json:"last_firmware_update" bson:"last_firmware_update"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
