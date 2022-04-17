package runner

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"os"
	"path"
)

func BuildConfig(players []database.PlayerVersion) Config {
	c := Config{
		Server:  config.Configuration.ServerCommand,
		Players: map[string]string{},
		Timeout: config.Configuration.PlayerTimeout,
	}

	for _, player := range players {
		c.Players[player.Player.Name] = player.Entrypoint
	}

	return c
}

func ProcessGame(game database.Game) error {
	game.Gamefolder = path.Join(config.Configuration.DataFolder, fmt.Sprintf("game-%06d", game.ID))
	database.Db.Save(&game)

	players := []string{}
	for _, player := range game.Players {
		players = append(players, player.Player.Name)
	}

	probojGame := Game{
		Gamefolder: game.Gamefolder,
		Players:    players,
		Args:       game.Map.Args,
	}

	probojConfig := BuildConfig(game.Players)

	err := RunGame(probojConfig, probojGame)
	if err != nil {
		game.State = database.GameDNF
		database.Db.Save(&game)
		return err
	}

	game.State = database.GameWaiting
	scores, err := ScoresFromFile(probojGame)
	if err != nil {
		return err
	}

	scoresString, err := scores.String()
	if err != nil {
		return err
	}
	game.Scores = scoresString
	database.Db.Save(&game)

	for playerName, score := range scores {
		database.Db.Model(&database.Player{}).Where("name = ?", playerName).Update("score", gorm.Expr("score + ?", score))
	}

	return nil
}

func ScoresFromFile(game Game) (database.Scores, error) {
	file, err := os.OpenFile(path.Join(game.Gamefolder, "score"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scores database.Scores
	d := json.NewDecoder(file)
	err = d.Decode(&scores)
	return scores, err
}
