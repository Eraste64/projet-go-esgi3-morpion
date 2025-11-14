package main

import (
	"projet-go-esgi3-morpion/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	defer db.Close()

	r := gin.Default() // Crée un serveur avec logs et récupération d'erreurs

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
