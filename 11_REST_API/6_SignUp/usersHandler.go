package apis

import (
	"net/http"
	"rest/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Can't Bind the Data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Can't save the User Data!!!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "New User Created Successfully..."})

}

func getAllUsers(context *gin.Context) {

	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can't fetch users Data"})
		return
	}

	context.JSON(http.StatusOK, users)

}
