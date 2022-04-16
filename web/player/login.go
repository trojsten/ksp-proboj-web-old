package player

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/database"
	"net/http"
	"strings"
)

func failed(c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic realm=\"Proboj management interface\"")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func AuthRequired(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" || !strings.HasPrefix(header, "Basic ") {
		failed(c)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(header, "Basic "))
	if err != nil {
		failed(c)
		return
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	username, password := parts[0], parts[1]

	var user database.Player
	database.Db.Where("name = ? AND password = ?", username, password).Limit(1).Find(&user)
	if user.ID == 0 {
		failed(c)
		return
	}

	c.Set("PROBOJ_PLAYER", user)
}
