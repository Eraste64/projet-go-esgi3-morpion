package database

import (
	"database/sql" // package standard pour manipuler une base SQL
	"log"          // package pour afficher les erreurs

	_ "modernc.org/sqlite" // driver SQLite, l'underscore active le driver mais on n'utilise pas directement le package
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite", "./morpion.db")
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT UNIQUE,
    password TEXT,
    isAdmin BOOLEAN
);
`

	createGamesTable := `
CREATE TABLE IF NOT EXISTS games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    grid TEXT,
    player TEXT,
    status TEXT
);
`

	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createGamesTable)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}
