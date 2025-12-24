package models

import "time"

type Events struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int       `json:"userId"`
}

var events = []Events{}

func (e Events) Save() {

	events = append(events, e)

}

func GetAllEvents() []Events {
	return events
}
