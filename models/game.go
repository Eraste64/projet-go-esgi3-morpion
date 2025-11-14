package models

type Game struct {
	ID       int          `json:"id"`
	Player1  string       `json:"player1"`
	Player2  string       `json:"player2"`
	Board    [3][3]string `json:"board"`    // vide au départ
	NextTurn string       `json:"nextTurn"` // Player1 ou Player2
	Winner   string       `json:"winner"`   // vide si pas de gagnant
}
