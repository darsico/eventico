package main

import (
	"net/http"
	"strconv"

	"example.com/eventico/db"
	"example.com/eventico/models"
	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
  server := gin.Default()

  server.GET("/events", getEvents)
  server.POST("/events", createEvent)
  server.GET("/events/:id", getEvent)

  server.Run(":8080") // localhost:8080
}

func getEvent(context *gin.Context) {
  rawEventId := context.Param("id")

  eventId, err := strconv.ParseInt(rawEventId, 10, 64)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "message": err.Error()})
    return
  }

  event, err := models.GetEventById(eventId)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{
      "message": "Error getting event",
      "error": err.Error()})
    return
  } 
  
  context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
  events, err := models.GetAllEvents()
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
  event := models.Event{}
  err := context.ShouldBindJSON(&event)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error":"Invalid data", "message": err.Error()})
    return
  }

  event.ID = 1
  event.UserID = 1
  err = event.Save()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating event", "message": err.Error()})
    return 
  }
  
  context.JSON(http.StatusCreated, gin.H{
    "message": "Event created successfully",
    "event": event,
  })
}