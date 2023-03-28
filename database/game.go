package database

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type GameState string

const (
	GameCreated GameState = "CREATED"
	GameRunning           = "RUNNING"
	GameWaiting           = "WAITING"
	GamePlaying           = "PLAYING"
	GameDone              = "DONE"
	GameDNF               = "DNF"
)

type Game struct {
	ID         uint `gorm:"primaryKey"`
	Gamefolder string
	Map        Map
	MapID      uint
	State      GameState
	Players    []PlayerVersion `gorm:"many2many:game_players;"`
	Scores     string
	CreatedAt  time.Time
}

type Scores map[string]int

func (g Game) Scoresf() (string, error) {
	var scores Scores
	err := json.Unmarshal([]byte(g.Scores), &scores)
	if err != nil {
		return "", err
	}

	var data []string
	for s, i := range scores {
		data = append(data, fmt.Sprintf("%s: %db", s, i))
	}

	return strings.Join(data, ", "), nil
}

func (s Scores) String() (string, error) {
	data, err := json.Marshal(s)
	return string(data), err
}
