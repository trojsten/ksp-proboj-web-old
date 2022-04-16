package player

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/proboj/web/compiler"
	"ksp.sk/proboj/web/config"
	"ksp.sk/proboj/web/database"
	"ksp.sk/proboj/web/web/utils"
	"net/http"
	"os"
	"path"
	"strconv"
)

func PostUpdate(c *gin.Context) {
	p, _ := c.Get("PROBOJ_PLAYER")
	player := p.(database.Player)

	fileHeader, err := c.FormFile("file")
	if err == http.ErrMissingFile {
		c.String(http.StatusBadRequest, "no file attached")
		return
	}
	if err != nil {
		utils.RenderError(c, "formfile", err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		utils.RenderError(c, "open file", err)
		return
	}

	lastVersion := 0
	database.Db.Raw("SELECT MAX(version) FROM player_versions WHERE player_id = ?", player.ID).Scan(&lastVersion)

	version := database.PlayerVersion{
		Player:     player,
		Version:    lastVersion + 1,
		Entrypoint: "",
		IsLatest:   false,
	}
	database.Db.Save(&version)

	root := path.Join(config.Configuration.UploadFolder, player.Name, strconv.Itoa(version.Version))
	err = os.MkdirAll(root, 0777)
	if err != nil {
		utils.RenderError(c, "create directory", err)
		return
	}

	err = compiler.Unpack(file, root)
	if err != nil {
		utils.RenderError(c, "unpack", err)
		return
	}

	compilerStderr, err := compiler.Compile(root)
	if compilerStderr != "" {
		c.String(500, "error while compiling:\n%s", compilerStderr)
		return
	}
	if err != nil {
		utils.RenderError(c, "compile", err)
		return
	}

	version.IsLatest = true
	version.Entrypoint = path.Join(root, "player")
	database.Db.Save(&version)

	if _, noredir := c.GetQuery("noredir"); noredir {
		c.String(200, "OK")
	} else {
		c.Redirect(http.StatusFound, "/management/")
	}
}
