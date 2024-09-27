package routes

import (
	"net/http"

	"example.com/eventico/models"
	"example.com/eventico/utils"
	"github.com/gin-gonic/gin"
)

func signup (context *gin.Context) {  
  user := models.User{}
  err := context.ShouldBindJSON(&user)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "message": err.Error()})
    return
  }

  err = user.Save()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user", "message": err.Error()})
    return
  }

  context.JSON(http.StatusCreated, gin.H{
    "message": "User created successfully",
  })
}

func login (context *gin.Context) {
  user := models.User{}

  err := context.ShouldBindJSON(&user)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "message": err.Error()})
    return
  }

  err = user.ValidateCredentials()

  if err != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }

  token, err := utils.GenerateToken(user.Email, user.ID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't authenticate user", 
      "error": err.Error()})
  
    return
  }

  context.JSON(http.StatusOK, gin.H{
    "message": "User logged in successfully",
    "token": token,
  })  
}