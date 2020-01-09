package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createDevice(c *gin.Context) {
	var device Device
	err := json.NewDecoder(c.Request.Body).Decode(&device)
	checkError(err, c)
	insertResult, err := DB.Collection.InsertOne(context.TODO(), device)
	checkError(err, c)
	device.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, device)
}
