package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var db = Database{Collection: connect()}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/devices/:id", readDevice)

	r.POST("/devices", createDevice)

	r.PUT("/devices/:id", updateDevice)

	r.DELETE("/devices/:id", deleteDevice)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
