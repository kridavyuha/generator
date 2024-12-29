package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BattingRules struct {
	Run             int `json:"run"`
	DotBall         int `json:"dot_ball"`
	BoundaryBonus   int `json:"boundary_bonus"`
	SixBonus        int `json:"six_bonus"`
	ThirtyRunBonus  int `json:"thirty_run_bonus"`
	FiftyRunBonus   int `json:"fifty_run_bonus"`
	HundredRunBonus int `json:"hundred_run_bonus"`
	GoldenDuck      int `json:"golden_duck"`
}

type BowlingRules struct {
	Wicket                       int `json:"wicket"`
	DotBall                      int `json:"dot_ball"`
	MaidenOver                   int `json:"maiden_over"`
	ThreeWicketBonus             int `json:"three_wicket_bonus"`
	FiveWicketBonus              int `json:"five_wicket_bonus"`
	EconomyRateBonusLessThanFour int `json:"economy_rate_bonus_less_than_four"`
	EconomyRateBonusLessThanSix  int `json:"economy_rate_bonus_less_than_six"`
}

type FieldingRules struct {
	Catch    int `json:"catch"`
	Stumping int `json:"stumping"`
	RunOut   int `json:"run_out"`
}

type FantasyRules struct {
	BattingRules  BattingRules  `json:"batting_rules"`
	BowlingRules  BowlingRules  `json:"bowling_rules"`
	FieldingRules FieldingRules `json:"fielding_rules"`
}

type FantasyCalc struct {
	RulesFile string `json:"rules_file"`
	TeamFile  string `json:"team_file"`
}

type Player struct {
	PlayerID string `json:"player_id"`
	Points   int    `json:"points"`
}

func (fc *FantasyCalc) NewFantasyCalc() *FantasyCalc {
	return &FantasyCalc{
		RulesFile: "rules.json",
		TeamFile:  "team.json",
	}
}

func (fc *FantasyCalc) CalculatePoints(ballDetails *BallData) map[string]int {

	// Load rules from file
	rules, err := loadRulesFromFileIntoMap(fc.RulesFile)
	if err != nil {
		return nil
	}

	// TODO: At later point use a set over here as we don't want to repeat the same player
	// i.e same player may get points for multiple events in a single ball
	points := make(map[string]int)
	// Three types of points...
	// 1. Wicket
	if ballDetails.Wicket == "1" {
		points[ballDetails.BowlerID] += rules.BowlingRules.Wicket
		if ballDetails.Method == "caught" {
			points[ballDetails.CaughtByID] += rules.FieldingRules.Catch
		}
		//TODO: also include run out and stumping
	}
	// 2. Runs
	if ballDetails.RunsFromBall != "0" {
		points[ballDetails.BatterID] += rules.BattingRules.Run * (int)(ballDetails.RunsFromBall[0]-'0')
		if ballDetails.RunsFromBall == "4" {
			points[ballDetails.BatterID] += rules.BattingRules.BoundaryBonus
		}
		if ballDetails.RunsFromBall == "6" {
			points[ballDetails.BatterID] += rules.BattingRules.SixBonus
		}
		//Todo: handler 30,50,100 run bonus...
	} else {
		points[ballDetails.BatterID] += rules.BattingRules.DotBall
	}
	// 3.Bowling Dots
	if ballDetails.RunsFromBall == "0" {
		points[ballDetails.BowlerID] += rules.BowlingRules.DotBall
	}

	fmt.Println(points)

	return points
}

func loadRulesFromFileIntoMap(rulesFile string) (*FantasyRules, error) {
	data, err := os.ReadFile(rulesFile)

	if err != nil {
		return nil, err
	}

	var rules *FantasyRules

	err = json.Unmarshal(data, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
