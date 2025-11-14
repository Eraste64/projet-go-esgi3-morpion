package main

import (
	"Projet-Go-ESGI3-Morpion/database"
)

func main() {
	db := database.Connect()
	defer db.Close()
}
