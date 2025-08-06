package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, gameState := CreateNewSession()

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

func getGameStateFromRequest(r *http.Request) (string, GameStateDB, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", GameStateDB{}, err
	}

	sessionID := cookie.Value
	gameState, err := GetGameState(sessionID)
	if err != nil {
		return "", GameStateDB{}, err
	}

	return sessionID, gameState, nil
}

func drawHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, gameState, err := getGameStateFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	gameState.Hand = gameState.Hand.Draw()
	prizeValue := gameState.Hand.GetPrizeValue(1) // Assuming a bet of 1 for now
	gameState.Credits += prizeValue

	err = UpdateGameState(sessionID, gameState)
	if err != nil {
		http.Error(w, "Failed to update game state", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

type HoldRequest struct {
	Cards []int `json:"cards"`
}

func holdHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, gameState, err := getGameStateFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	var req HoldRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := 0; i < 5; i++ {
		held := false
		for _, cardIndex := range req.Cards {
			if i == cardIndex {
				held = true
				break
			}
		}
		if gameState.Hand.IsHeld(i) != held {
			gameState.Hand = gameState.Hand.ToggleHold(i)
		}
	}

	err = UpdateGameState(sessionID, gameState)
	if err != nil {
		http.Error(w, "Failed to update game state", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

func StartAPIServer() {
	http.HandleFunc("/api/game/new", newGameHandler)
	http.HandleFunc("/api/game/draw", drawHandler)
	http.HandleFunc("/api/game/hold", holdHandler)
	http.ListenAndServe(":8080", nil)
}
