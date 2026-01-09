package apis

import "github.com/gin-gonic/gin"

func RegisterRouters(server *gin.Engine) {

	server.GET("/events", getEvents)          //[GET] to get all event's
	server.GET("/events/:id", getEventById)   //[GET] to get event based on ID
	server.POST("/createEvent", createEvent)  //[POST] Create Event
	server.PUT("/events/:id", updateEvent)    //[PUT] update the event
	server.DELETE("/events/:id", deleteEvent) //[DELETE] delete the event
	server.POST("/signup", signup)
	server.GET("/users")

}
