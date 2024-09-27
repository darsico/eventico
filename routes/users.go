package routes

import (
	"net/http"

	"example.com/eventico/models"
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

  context.JSON(http.StatusOK, gin.H{
    "message": "User logged in successfully",
  })  
}