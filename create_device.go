package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createDevice(c *gin.Context) {
	var device Device
	err := json.NewDecoder(c.Request.Body).Decode(&device)
	checkError(err, c)
	device.Serial, err = validation.EnsureString(&device.Serial)
	checkError(err, c)
	device.Name, err = validation.EnsureString(&device.Name)
	checkError(err, c)
	insertResult, err := DB.Collection.InsertOne(context.TODO(), device)
	checkError(err, c)
	device.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, device)
}
