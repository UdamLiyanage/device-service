package main

import "github.com/gin-gonic/gin"

func createDevice(c *gin.Context) {
	c.String(200, "Create Device Function")
}
