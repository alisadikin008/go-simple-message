package main

import (
	messageCon "message-api/message"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	version := router.Group("/api/v1")
	{
		version.POST("/messages/", messageCon.PostMessage)     // post new message
		version.GET("/messages/", messageCon.AllMessage)       // get all message (including previous message)
		version.GET("/messages/realtime", messageCon.RealTime) // get realtime message
	}

	router.Run()
}
