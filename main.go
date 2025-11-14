package main

import (
	"projet-go-esgi3-morpion/database"
)

func main() {
	db := database.Connect()
	defer db.Close()
}
