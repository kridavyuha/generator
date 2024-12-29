package main

// Get fixtures from the API
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fixture struct
type Fixture struct {
	MatchID string `json:"match_id"`
	TeamA   string `json:"team_a"`
	TeamB   string `json:"team_b"`
}

type PlayingTeams struct {
	TeamA string `json:"team_a"`
	TeamB string `json:"team_b"`
}

func (app *App) GetFixtures(w http.ResponseWriter, r *http.Request) {

	// Get match_id from request
	matchID := r.URL.Query().Get("match_id")
	if matchID == "" {
		http.Error(w, "match_id is required", http.StatusBadRequest)
		return
	}

	// Read fixtures.json file
	data, err := ioutil.ReadFile("/Users/rithvik/Documents/Strategic Fantasy League/Generator/fixtures.json")
	if err != nil {
		http.Error(w, "Unable to read fixtures file", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON data into fixtures slice
	var fixtures []Fixture
	err = json.Unmarshal(data, &fixtures)
	if err != nil {
		http.Error(w, "Unable to parse fixtures data", http.StatusInternalServerError)
		return
	}

	// Filter fixtures by match_id
	var fixtureFound *Fixture
	for _, fixture := range fixtures {
		if fixture.MatchID == matchID {
			fixtureFound = &fixture
			break
		}
	}

	if fixtureFound == nil {
		http.Error(w, "No fixture found with the given match_id", http.StatusNotFound)
		return
	}

	// Write fixture to response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fixtureFound)
	if err != nil {
		http.Error(w, "Unable to encode fixture data", http.StatusInternalServerError)
		return
	}
}
