package runner

import (
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"log"
	"time"
)

func RunnerLoop() {
	for {
		runnerTick()
		time.Sleep(15 * time.Second)
	}
}

func runnerTick() {
	var game database.Game
	database.Db.Where("state = ?", database.GameCreated).Order("id asc").Limit(1).
		Preload("Map").Preload("Players").Preload("Players.Player").Find(&game)
	if game.ID == 0 {
		return
	}

	log.Printf("Found game %d\n", game.ID)
	err := ProcessGame(game)
	if err != nil {
		log.Printf("Error while running game %d: %s\n", game.ID, err.Error())
		return
	}
}

func GeneratorLoop() {
	for {
		generatorTick()
		time.Sleep(5 * time.Second)
	}
}

func generatorTick() {
	var pendingGames int64
	database.Db.Model(&database.Game{}).Where("state = ? OR state = ?", database.GameCreated, database.GameWaiting).Count(&pendingGames)
	if pendingGames >= int64(config.Configuration.GamesAhead) {
		return
	}

	log.Println("Generating new game...")
	game, err := GenerateGame()
	if err != nil {
		log.Printf("Error while generating new game: %s\n", err.Error())
		return
	}
	log.Printf("Successfully generated game %d.\n", game.ID)
}
