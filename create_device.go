package main

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func createDevice(c echo.Context) error {
	device := new(Device)
	if err := c.Bind(device); err != nil {
		return err
	}
	err := validation.Validate(&device)
	if checkError(err) {
		return c.JSON(422, err)
	}
	device.CreatedAt, device.UpdatedAt = time.Now(), time.Now()
	insertResult, err := collection.InsertOne(context.TODO(), device)
	if checkError(err) {
		return c.JSON(500, err)
	}
	device.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.JSON(201, device)
}
