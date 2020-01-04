package main

import "github.com/gin-gonic/gin"

func deleteDevice(c *gin.Context) {
	c.String(200, "Delete Device Function")
}
