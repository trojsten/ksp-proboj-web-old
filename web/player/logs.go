package player

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"net/http"
	"os"
	"path"
)

func GetLog(c *gin.Context) {
	p, _ := c.Get("PROBOJ_PLAYER")
	player := p.(database.Player)

	var game database.Game
	database.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&game)

	if game.ID == 0 || (game.State != database.GameDone && game.State != database.GamePlaying) {
		c.String(404, "not found")
		return
	}

	root := path.Join(config.Configuration.DataFolder, fmt.Sprintf("game-%06d", game.ID), "logs", fmt.Sprintf("%s.gz", player.Name))
	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			c.String(http.StatusNotFound, "The log file was not found.")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.File(root)
}

func GetServerLog(c *gin.Context) {
	var game database.Game
	database.Db.Where("id = ?", c.Param("id")).Limit(1).Find(&game)

	if game.ID == 0 || (game.State != database.GameDone && game.State != database.GamePlaying) {
		c.String(404, "not found")
		return
	}

	root := path.Join(config.Configuration.DataFolder, fmt.Sprintf("game-%06d", game.ID), "logs", "__server.gz")
	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			c.String(http.StatusNotFound, "The log file was not found.")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.File(root)
}
