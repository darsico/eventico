package main

import (
	"example.com/eventico/db"
	"example.com/eventico/routes"
	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
  server := gin.Default()

  routes.RegisterRoutes(server)

  server.Run(":8080") // localhost:8080
}
