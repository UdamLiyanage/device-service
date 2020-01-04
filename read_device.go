package main

import "github.com/gin-gonic/gin"

func readDevice(c *gin.Context) {
	c.String(200, "Read Device Function")
}
