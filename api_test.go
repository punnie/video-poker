package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGameHandler(t *testing.T) {
	InitDB()

	req, err := http.NewRequest("POST", "/api/game/new", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(newGameHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check for session cookie
	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "session_id" {
		t.Errorf("Expected cookie name 'session_id', got '%s'", cookies[0].Name)
	}

	// Check the response body
	var gameState GameStateDB
	err = json.NewDecoder(rr.Body).Decode(&gameState)
	if err != nil {
		t.Fatal(err)
	}

	if gameState.Credits != 100 {
		t.Errorf("Expected initial credits to be 100, got %d", gameState.Credits)
	}

	if len(gameState.Hand.Hand.Cards) != 5 {
		t.Errorf("Expected hand to have 5 cards, got %d", len(gameState.Hand.Hand.Cards))
	}
}
