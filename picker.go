package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/rand"
)

type BallPicker struct {
	SummaryFile string `json:"summary_file"`
	Offset      int    `json:"offset"`
	MaxDelay    int    `json:"max_delay"`
	FantasyCalc *FantasyCalc
}

type BallData struct {
	BallNo       int    `json:"ballNo"`
	MatchID      string `json:"matchId"`
	Innings      string `json:"innings"`
	Over         string `json:"over"`
	Ball         string `json:"ball"`
	Batter       string `json:"batter"`
	BatterID     string `json:"batterId"`
	NonStriker   string `json:"nonStriker"`
	NonStrikerID string `json:"nonStrikerId"`
	Bowler       string `json:"bowler"`
	BowlerID     string `json:"bowlerId"`
	BatterRuns   string `json:"batterRuns"`
	ExtraRuns    string `json:"extraRuns"`
	RunsFromBall string `json:"runsFromBall"`
	Wicket       string `json:"wicket"`
	Method       string `json:"method"`
	PlayerOut    string `json:"playerOut"`
	PlayerOutID  string `json:"playerOutId"`
	CaughtBy     string `json:"caughtBy"`
	CaughtByID   string `json:"caughtById"`
}

type PerfDetails struct {
	PerfFactor  map[string]int `json:"perf_factor"`
	MatchID     string         `json:"match_id"`
	IsCompleted bool           `json:"isCompleted"`
}

func (bp *BallPicker) NewBallPicker() *BallPicker {
	return &BallPicker{
		SummaryFile: "./data/ball_by_ball_ipl.json",
		Offset:      1,
		MaxDelay:    30,
		FantasyCalc: (&FantasyCalc{}).NewFantasyCalc(),
	}
}

func (bp *BallPicker) StartMatch(ch *amqp.Channel) {
	// Load balls from file
	balls, err := loadBallsIntoMap(bp.SummaryFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ch.ExchangeDeclare(
		"balls",  // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Loop through balls
	for i := bp.Offset; i <= len(balls); i++ {
		// rand Sleep
		time.Sleep(time.Duration(rand.Intn(bp.MaxDelay)) * time.Second)
		// Get ball data
		ball, ok := balls[i]
		if !ok {
			fmt.Printf("Ball %d not found\n", i)
			continue
		}
		// Print ball data
		fmt.Printf("Ball %d: %s to %s, %s runs\n", ball.BallNo, ball.Bowler, ball.Batter, ball.RunsFromBall)
		ballByBall := PerfDetails{
			PerfFactor:  bp.FantasyCalc.CalculatePoints(&ball),
			MatchID:     ball.MatchID,
			IsCompleted: false,
		}
		fmt.Println(ballByBall)

		// respCode := PostRequest("http://localhost:8080/points", ballByBall)
		// fmt.Printf("Response code: %d\n", respCode)

		// Publish message
		body, err := json.Marshal(ballByBall)
		if err != nil {
			fmt.Printf("Error marshalling ballByBall: %v\n", err)
			continue
		}

		err = ch.PublishWithContext(ctx,
			"balls", // exchange
			"",      // routing key
			false,   // mandatory
			false,   // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		//TODO: Handler Error

	}

	PostRequest("http://localhost:8080/points", PerfDetails{
		IsCompleted: true,
	})
}

func loadBallsIntoMap(filename string) (map[int]BallData, error) {
	// Read JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Parse JSON array
	var balls []BallData
	if err := json.Unmarshal(data, &balls); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Create map with ball number as key
	ballMap := make(map[int]BallData)
	for _, ball := range balls {
		ballMap[ball.BallNo] = ball
	}

	return ballMap, nil
}
