package public

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/web/utils"
)

func GetIndex(c *gin.Context) {
	utils.Render(c, "index.gohtml", gin.H{})
}
