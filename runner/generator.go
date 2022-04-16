package runner

import (
	"fmt"
	"ksp.sk/proboj/web/database"
	"math/rand"
)

func GenerateGame() (database.Game, error) {
	var versions []database.PlayerVersion
	database.Db.Where("is_latest = 1").Find(&versions)
	if len(versions) < 2 {
		return database.Game{}, fmt.Errorf("not enough players available")
	}

	var maps []database.Map
	database.Db.Find(&maps)
	if len(versions) == 0 {
		return database.Game{}, fmt.Errorf("no available maps")
	}

	rand.Shuffle(len(versions), func(i, j int) {
		versions[i], versions[j] = versions[j], versions[i]
	})
	pickedMap := maps[rand.Intn(len(maps))]

	game := database.Game{
		Map:     pickedMap,
		State:   database.GameCreated,
		Players: versions,
	}

	database.Db.Save(&game)
	return game, nil
}
