package main

import (
	"encoding/json"
	"net/http"
)

func (app *App) GetScores(w http.ResponseWriter, r *http.Request) {
	// Create a new picker
	bp := (&BallPicker{}).NewBallPicker()

	// Start the match
	go bp.StartMatch()
	// Write scores to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]byte("Match started"))
}
