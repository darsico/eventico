package main

import (
	"net/http"

	"example.com/eventico/db"
	"example.com/eventico/models"
	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
  server := gin.Default()

  server.GET("/events", getEvents)
  server.POST("/events", createEvent)

  server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
  events := models.GetAllEvents()
  context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
  event := models.Event{}
  err := context.ShouldBindJSON(&event)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  event.ID = 1
  event.UserID = 1
  event.Save()
  context.JSON(http.StatusCreated, gin.H{
    "message": "Event created successfully",
    "event": event,
  })
}