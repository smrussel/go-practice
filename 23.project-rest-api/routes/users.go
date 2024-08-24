package routes

import (
	"log"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user.",
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.ValidateCredentials()

	log.Println("err validate==>", err)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user*.",
		})
		return
	}

	log.Println("user ===>", user)

	token, err := utils.GenerateToken(user.Email, user.ID)

	log.Println(token)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user**.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successfull.",
		"token":   token,
	})
}
