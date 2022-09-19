package main

import (
 "github.com/gin-gonic/gin"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/users/:auth", getUsers)
	r.POST("/user/:auth", addUser)
	
	r.Run(":8001")
}