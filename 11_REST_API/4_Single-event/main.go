package main

import (
	"net/http"
	"rest/db"
	"rest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)
	server.POST("/createEvent", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "msg": "can't fetch data from DATA-BASE!"})
		return
	}
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {

	var event models.Events

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "msg": "Somthing Wrong"})
		return
	}

	event.Id = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "msg": "Could  Not Create Event!!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"msg": "Created Successfully", "Events": event})

}

func getEventById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not convert to Int64"})
		return
	}

	event, err := models.GetEventFromId(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch DATA from invalid ID"})
		return
	}

	context.JSON(http.StatusOK, event)

}
