package public

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
	"os"
	"path"
)

func GetGames(c *gin.Context) {
	var games []database.Game
	database.Db.Preload("Players").Preload("Players.Player").Preload("Map").Order("created_at desc").Find(&games)

	utils.Render(c, "games.gohtml", gin.H{
		"games": games,
	})
}

func GetObserverLog(c *gin.Context) {
	var game database.Game
	database.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&game)

	if game.ID == 0 || (game.State != database.GameDone && game.State != database.GamePlaying) {
		c.String(404, "not found")
		return
	}

	file := path.Join(game.Gamefolder, "observer.gz")
	ext := ".gz"
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		file = path.Join(game.Gamefolder, "observer")
		ext = ""
	}

	c.FileAttachment(file, fmt.Sprintf("observer-%06d%s", game.ID, ext))
}
