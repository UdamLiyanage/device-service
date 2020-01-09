package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateDevice(c *gin.Context) {
	var device Device
	err := json.NewDecoder(c.Request.Body).Decode(&device)
	checkError(err, c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name": device.Name,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.JSON(200, res)
}
