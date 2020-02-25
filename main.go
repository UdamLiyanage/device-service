package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	auth := gin.BasicAuth(gin.Accounts{
		os.Getenv("API_AUTH_USERNAME"): os.Getenv("API_AUTH_PASSWORD"),
	})
	r.Use(auth)

	r.GET("/devices/:id", readDevice)

	r.POST("/devices", createDevice)

	r.PUT("/devices/:id", updateDevice)

	r.DELETE("/devices/:id", deleteDevice)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":8002"))
}
