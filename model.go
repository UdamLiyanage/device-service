package main

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Paginator interface {
		Paginate() PaginateResult
		PageCheck(c echo.Context) error
	}

	Device struct {
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

	Firmware struct {
		ID         primitive.ObjectID `json:"_id" bson:"_id"`
		Version    string             `json:"firmware_version" bson:"firmware_version"`
		LastUpdate time.Time          `json:"last_firmware_update_date" bson:"last_firmware_update"`
	}

	Pager struct {
		Limit     int64
		URL       string
		FirstPage bool
		Filter    bson.D
		QueryVal  string
		QueryKey  string
	}

	PaginateResult struct {
		NextPage string   `json:"next_page"`
		PrevPage string   `json:"previous_page"`
		Data     []Device `json:"data"`
	}
)
