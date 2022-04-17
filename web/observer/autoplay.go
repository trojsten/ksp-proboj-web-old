package observer

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
	"net"
)

func GetAutoPlay(c *gin.Context) {
	ip, _ := c.RemoteIP()
	presenter := net.ParseIP(config.Configuration.PresenterIP)
	if !ip.Equal(presenter) {
		c.String(403, "forbidden %s", ip.String())
		return
	}

	var game database.Game
	database.Db.Model(&database.Game{}).Where("state = ?", database.GamePlaying).Update("state", database.GameDone)
	database.Db.Where("state = ?", database.GameWaiting).Order("id asc").Limit(1).Find(&game)
	//game.State = database.GamePlaying
	//database.Db.Save(&game)

	var players []database.Player
	database.Db.Order("score desc").Find(&players)

	utils.Render(c, "_autoplay.gohtml", gin.H{
		"game":    game,
		"players": players,
	})
}
