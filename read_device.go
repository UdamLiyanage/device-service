package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readDevice(c *gin.Context) {
	var device Device
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objID}
	err = db.Collection.FindOne(context.TODO(), filter).Decode(&device)
	if err != nil {
		panic(err)
	}
	c.JSON(200, device)
}
