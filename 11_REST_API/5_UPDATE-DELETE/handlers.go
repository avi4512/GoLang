package apis

import (
	"net/http"
	"rest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func updateEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not convert to Int64"})
		return
	}

	_, err = models.GetEventFromId(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not UPDATE the DATA from invalid ID"})
		return
	}

	var updateEvent models.Events

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not convert to Int64"})
		return
	}

	updateEvent.Id = id

	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could UPDATE the DATA from invalid ID"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated Successfully...."})

}

func deleteEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the ID!"})
		return
	}

	events, err := models.GetEventFromId(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event id NOT found or INVALID ID!"})
		return
	}

	err = events.DeleteEventById()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event could not be DELETE'd !!!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event [DELETED] Successfully...."})

}
