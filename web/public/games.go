package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
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

	c.FileAttachment(path.Join(game.Gamefolder, "observer"), fmt.Sprintf("observer-%d", game.ID))
}
