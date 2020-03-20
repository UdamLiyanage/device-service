package main

import (
	"context"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func updateDevice(c echo.Context) error {
	var device Device
	err := json.NewDecoder(c.Request().Body).Decode(&device)
	if checkError(err) {
		return c.JSON(500, err)
	}
	err = validation.Validate(device.Name,
		validation.Required,
		validation.Length(5, 25),
	)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name": device.Name,
		},
	}
	device.UpdatedAt = time.Now()
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

func updateDeviceFirmware(c echo.Context) error {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"firmware_version":      c.Param("firmware_version"),
			"last_firmware_version": time.Now(),
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if checkError(err) {
		return err
	}
	return c.JSON(200, nil)
}
