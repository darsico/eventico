package routes

import (
	"net/http"
	"strconv"

	"example.com/eventico/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
  userId := context.GetInt64("userId")
  rawEventId := context.Param("id")
  eventId, err := strconv.ParseInt(rawEventId, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "message": err.Error()})
    return
  }
  
  event, err := models.GetEventById(eventId)

  if err != nil || event.ID == 0 {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting event", "error": err.Error()})
    return
  }

  err = event.Register(userId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event", "error": err.Error()})
    return
  }

  context.JSON(http.StatusCreated,gin.H{
      "message": "Successfully registered for event", 
      "event": event,
    })
}

func cancelRegistration(context *gin.Context) {
  userId := context.GetInt64("userId")
  rawEventId := context.Param("id")
  eventId, err := strconv.ParseInt(rawEventId, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "message": err.Error()})
    return
  }

  var event models.Event

  event.ID = eventId

  err = event.CancelRegistration(userId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Error cancelling registration", "error": err.Error()})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration"})
} 