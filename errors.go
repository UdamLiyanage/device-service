package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkError(err error, c *gin.Context) {
	if err == nil {
		return
	}
	if err == mongo.ErrNoDocuments {
		c.AbortWithStatusJSON(404, err)
	} else {
		c.AbortWithStatusJSON(500, err)
	}
}
