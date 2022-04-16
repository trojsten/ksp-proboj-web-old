package player

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
)

func GetPlayerSite(c *gin.Context) {
	p, _ := c.Get("PROBOJ_PLAYER")
	player := p.(database.Player)

	var versions []database.PlayerVersion
	database.Db.Where("player_id = ?", player.ID).Order("version desc").Find(&versions)

	utils.Render(c, "mgmt_index.gohtml", gin.H{
		"player":   player,
		"versions": versions,
	})
}
