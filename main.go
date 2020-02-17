package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:  []string{"auth.cliko.io"},
		SSLRedirect:   true,
		STSSeconds:    31536000,
		FrameDeny:     true,
		IsDevelopment: false,
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			if err != nil {
				c.Abort()
				return
			}

			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	r := gin.Default()
	auth := gin.BasicAuth(gin.Accounts{
		"udam": "g4}U2)$S_q7n=aH#2WRj",
	})
	r.Use(auth)
	r.Use(secureFunc)

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
