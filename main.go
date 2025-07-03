package main

import (
	"CRUD_DOCKER_HOMEPAGE/db"
	"CRUD_DOCKER_HOMEPAGE/handle"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	handle.Testrsa()
	router.POST("/login_1", handle.Login)
	router.POST("/testtoken", handle.TestToken)

	router.POST("/receive_formdata", handle.Receiveinfo_useformdata)

	router.POST("/receive_rawdata", handle.Receiveinfo_userawdata)

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
