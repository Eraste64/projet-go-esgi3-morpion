package main

import (
	"projet-go-esgi3-morpion/database"
	"projet-go-esgi3-morpion/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	defer db.Close()

	r := gin.Default() // Crée un serveur avec logs et récupération d'erreurs

	// Endpoint de test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Créer un nouvel utilisateur
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

	// Afficher tous les utilisateurs
	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT ID, Name FROM users")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			rows.Scan(&user.ID, &user.Name)
			users = append(users, user)
		}

		c.JSON(200, users)
	})

	// Afficher un utilisateur précis
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id") // Récupère l'ID depuis l'URL

		var user models.User
		err := db.QueryRow("SELECT ID, Name FROM users WHERE ID = ?", id).Scan(&user.ID, &user.Name)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	})

	// Mettre à jour le pseudo d'un utilisateur
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id") // Récupère l'ID depuis l'URL

		// Convertit l'ID en int
		numID, errConv := strconv.Atoi(id)
		if errConv != nil {
			c.JSON(400, gin.H{"error": "ID invalide"})
			return
		}

		var updatedUser models.User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Vérifie si le pseudo existe déjà pour un autre utilisateur
		var existingID int
		err := db.QueryRow("SELECT ID FROM users WHERE Name = ?", updatedUser.Name).Scan(&existingID)
		if err == nil && existingID != numID { // pseudo déjà utilisé par un autre utilisateur
			c.JSON(400, gin.H{"error": "Pseudo déjà utilisé"})
			return
		}

		// Met à jour le pseudo dans la BDD
		_, err = db.Exec("UPDATE users SET Name = ? WHERE ID = ?", updatedUser.Name, numID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Pseudo mis à jour", "id": numID, "name": updatedUser.Name})
	})

	// Supprimer un utilisateur
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Convertit l'ID en int
		numID, errConv := strconv.Atoi(id)
		if errConv != nil {
			c.JSON(400, gin.H{"error": "ID invalide"})
			return
		}

		// Supprime l'utilisateur dans la BDD
		res, err := db.Exec("DELETE FROM users WHERE ID = ?", numID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Vérifie si un utilisateur a été supprimé
		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"error": "Utilisateur non trouvé"})
			return
		}

		c.JSON(200, gin.H{"message": "Utilisateur supprimé", "id": numID})
	})

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
