package public

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/web/utils"
)

func GetDocs(c *gin.Context) {
	utils.Render(c, "docs.gohtml", gin.H{})
}
