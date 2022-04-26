package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/runner"
	"ksp.sk/proboj/web/web/utils"
	"path"
	"strconv"
)

func GetGames(c *gin.Context) {
	var number int
	var err error
	numberString, present := c.GetQuery("n")
	if !present {
		number = 1
	} else {
		number, err = strconv.Atoi(numberString)
		if err != nil {
			utils.RenderError(c, "int convert", err)
			return
		}
	}

	games := []runner.Game{}
	for i := 0; i < number; i++ {
		gMap, err := runner.RandomMap()
		if err != nil {
			utils.RenderError(c, "random map", err)
			return
		}

		playerVers, err := runner.RandomPlayers(gMap)
		if err != nil {
			utils.RenderError(c, "random players", err)
			return
		}

		players := []string{}
		for _, ver := range playerVers {
			players = append(players, ver.Player.Name)
		}

		games = append(games, runner.Game{
			Gamefolder: path.Join("ext", fmt.Sprintf("game-%06d", i+1)),
			Players:    players,
			Args:       gMap.Args,
		})
	}

	c.JSON(200, games)
}

func GetConfig(c *gin.Context) {
	var players []database.PlayerVersion
	database.Db.Where("is_latest = 1").Preload("Player").Find(&players)
	c.JSON(200, runner.BuildConfig(players))
}
