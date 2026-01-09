package main

import (
	"rest/apis"
	"rest/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	apis.RegisterRouters(server)

	server.Run(":8080")

}
