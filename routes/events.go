package routes

import (
	"net/http"
	"strconv"

	"example.com/eventico/models"
	"github.com/gin-gonic/gin"
)


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

  userId := context.GetInt64("userId")
  event.UserID = userId

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

func updateEvent (context *gin.Context) {
  rawEventId := context.Param("id")
  eventId, err := strconv.ParseInt(rawEventId, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "message": err.Error()})
    return
  }

  userId := context.GetInt64("userId")
  event, err := models.GetEventById(eventId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event", "message": err.Error()})
    return
  }

  if event.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to update this event"})
    return
  }

  var updatedEvent models.Event
  
  err = context.ShouldBindJSON(&updatedEvent)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "message": err.Error()})
    return
  }

  updatedEvent.ID = eventId

  err = updatedEvent.Update()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating event", "message": err.Error()})
    return
  }

  context.JSON(http.StatusOK, gin.H{
    "message": "Event updated successfully!",
  })
}

func deleteEvent(context *gin.Context) {
  rawEventId := context.Param("id")
  eventId, err := strconv.ParseInt(rawEventId, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "message": err.Error()})
    return
  }

  userId := context.GetInt64("userId")
  event, err := models.GetEventById(eventId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting event", "message": err.Error()})
    return
  }

  if event.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to delete this event"})
    return
  }

  err = event.Delete()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting event", "message": err.Error()})
    return
  }

  context.JSON(http.StatusOK, gin.H{
    "message": "Event deleted successfully!",
  })
} 