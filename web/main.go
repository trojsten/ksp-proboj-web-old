package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"ksp.sk/proboj/web/web/api"
	"ksp.sk/proboj/web/web/observer"
	"ksp.sk/proboj/web/web/player"
	"ksp.sk/proboj/web/web/public"
	"path"
	"strings"
	"time"
)
import "github.com/gin-contrib/multitemplate"

func prepareTamplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	dir, err := ioutil.ReadDir("templates")
	if err != nil {
		panic(err)
	}

	funcs := template.FuncMap{
		"formatDate": formatDate,
	}

	for _, info := range dir {
		if info.IsDir() {
			continue
		}

		if info.Name() == "base.gohtml" {
			continue
		}

		if strings.HasPrefix(info.Name(), "_") {
			r.AddFromFilesFuncs(info.Name(), funcs, path.Join("templates", info.Name()))
		} else {
			r.AddFromFilesFuncs(info.Name(), funcs, "templates/base.gohtml", path.Join("templates", info.Name()))
		}
	}
	return r
}

func formatDate(t time.Time) string {
	return t.Format("02.01. 15:04:05")
}

func Start() {
	r := gin.Default()
	r.HTMLRender = prepareTamplates()

	r.Static("static/", "static")
	r.Static("download/", "download")

	r.GET("/", public.GetDocs)
	r.GET("/scores/", public.GetScores)
	r.GET("/games/", public.GetGames)
	r.GET("/games/:id/observer", public.GetObserverLog)

	r.Static("/observer/", "observer")
	r.GET("/autoplay/", observer.GetAutoPlay)

	auth := r.Group("management", player.AuthRequired)
	auth.GET("/", player.GetPlayerSite)
	auth.POST("/update/", player.PostUpdate)
	auth.GET("/log/:id", player.GetLog)
	auth.GET("/log/:id/server", player.GetServerLog)

	r.GET("/api/games/", api.GetGames)
	r.GET("/api/config/", api.GetConfig)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
