package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Squad struct {
	Team    string          `json:"team"`
	Players []PlayerInSquad `json:"players"`
	Id      int             `json:"id"`
}

type PlayerInSquad struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (app *App) GetSquads(w http.ResponseWriter, r *http.Request) {
	// Get squads from the API
	teamName := r.URL.Query().Get("team_name")
	if teamName == "" {
		http.Error(w, "team_name is required", http.StatusBadRequest)
		return
	}

	file, err := os.ReadFile("/Users/rithvik/Documents/Strategic Fantasy League/Generator/squads.json")
	if err != nil {
		http.Error(w, "Unable to read squads file", http.StatusInternalServerError)
		return
	}

	var squads []Squad
	err = json.Unmarshal(file, &squads)
	if err != nil {
		http.Error(w, "Unable to parse squads data", http.StatusInternalServerError)
		return
	}

	var squadFound *Squad
	for _, squad := range squads {
		if squad.Team == teamName {
			squadFound = &squad
			break
		}
	}

	if squadFound == nil {
		http.Error(w, "No squad found with the given team_name", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(squadFound)

}
