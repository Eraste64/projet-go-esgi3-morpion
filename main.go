package main

import (
	"projet-go-esgi3-morpion/database"
	"projet-go-esgi3-morpion/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	defer db.Close()

	r := gin.Default() // Crée un serveur avec logs et récupération d'erreurs

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/users", func(c *gin.Context) {
		var newUser models.User

		// Lit le JSON envoyé par le client et le place dans la struct
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Insère l'utilisateur dans la BDD
		_, err := db.Exec("INSERT INTO users (Name) VALUES (?)", newUser.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Renvoie le nouvel utilisateur en JSON
		c.JSON(201, newUser)
	})

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
