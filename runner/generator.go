package runner

import (
	"fmt"
	"ksp.sk/proboj/web/database"
	"math/rand"
)

func RandomMap() (database.Map, error) {
	var maps []database.Map
	database.Db.Where("is_enabled = 1").Find(&maps)
	if len(maps) == 0 {
		return database.Map{}, fmt.Errorf("no available maps")
	}

	return maps[rand.Intn(len(maps))], nil
}

func RandomPlayers(mp database.Map) ([]database.PlayerVersion, error) {
	var versions []database.PlayerVersion
	database.Db.Where("is_latest = 1").Preload("Player").Find(&versions)
	if len(versions) < 2 {
		return versions, fmt.Errorf("not enough players available")
	}

	rand.Shuffle(len(versions), func(i, j int) {
		versions[i], versions[j] = versions[j], versions[i]
	})

	if len(versions) > mp.MaxPlayers {
		versions = versions[:mp.MaxPlayers]
	}

	return versions, nil
}

func GenerateGame() (database.Game, error) {
	pickedMap, err := RandomMap()
	if err != nil {
		return database.Game{}, err
	}

	versions, err := RandomPlayers(pickedMap)
	if err != nil {
		return database.Game{}, err
	}

	game := database.Game{
		Map:     pickedMap,
		State:   database.GameCreated,
		Players: versions,
	}

	database.Db.Save(&game)
	return game, nil
}
