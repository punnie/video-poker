package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/punnie/video-poker/pkg"
)

type GameStateDB struct {
	Credits int
	Hand    pkg.Hand
}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./poker.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		"id" TEXT NOT NULL PRIMARY KEY,
		"credits" INTEGER,
		"hand" TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func GetGameState(sessionID string) (GameStateDB, error) {
	var gameState GameStateDB
	var handJSON string

	err := db.QueryRow("SELECT credits, hand FROM sessions WHERE id = ?", sessionID).Scan(&gameState.Credits, &handJSON)
	if err != nil {
		return GameStateDB{}, err
	}

	err = json.Unmarshal([]byte(handJSON), &gameState.Hand)
	if err != nil {
		return GameStateDB{}, err
	}

	return gameState, nil
}

func CreateNewSession() (string, GameStateDB) {
	sessionID := uuid.New().String()
	gameState := GameStateDB{
		Credits: 100,
		Hand:    pkg.InitializeHand(),
	}

	handJSON, err := json.Marshal(gameState.Hand)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO sessions (id, credits, hand) VALUES (?, ?, ?)", sessionID, gameState.Credits, handJSON)
	if err != nil {
		log.Fatal(err)
	}

	return sessionID, gameState
}

func UpdateGameState(sessionID string, gameState GameStateDB) error {
	handJSON, err := json.Marshal(gameState.Hand)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE sessions SET credits = ?, hand = ? WHERE id = ?", gameState.Credits, handJSON, sessionID)
	return err
}
