package main

import (
	"apis/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/createEvent", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()
	context.JSON(200, events)
	//context.JSON(http.StatusOK, gin.H{"Meassage": "Hello Hii"})

}

func createEvent(context *gin.Context) {

	var event models.Events
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Meassage": err.Error()})
		return
	}

	event.Id = 1
	event.UserId = 1701
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"Meassage": "event Created", "event": event})

}
