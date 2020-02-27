package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readDevice(c echo.Context) error {
	var device Device
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	err = DB.Collection.FindOne(context.TODO(), filter).Decode(&device)
	if err != nil {
		return c.JSON(404, nil)
	}
	return c.JSON(200, device)
}

func readAllDevices(c echo.Context) error {
	var devices []Device
	userID := c.Param("id")
	filter := bson.D{{"uid", userID}}
	cur, err := DB.Collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var device Device
		err := cur.Decode(&device)
		if err != nil {
			panic(err)
		}
		devices = append(devices, device)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return c.JSON(200, devices)
}
