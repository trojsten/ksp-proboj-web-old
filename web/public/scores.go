package public

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
)

func GetScores(c *gin.Context) {
	var players []database.Player
	database.Db.Order("score desc").Find(&players)

	utils.Render(c, "scores.gohtml", gin.H{
		"players": players,
	})
}
