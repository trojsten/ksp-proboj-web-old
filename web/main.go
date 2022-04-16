package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"ksp.sk/proboj/web/web/player"
	"ksp.sk/proboj/web/web/public"
	"path"
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

		r.AddFromFilesFuncs(info.Name(), funcs, "templates/base.gohtml", path.Join("templates", info.Name()))
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

	r.GET("/", public.GetIndex)
	r.GET("/docs/", public.GetDocs)
	r.GET("/games/", public.GetGames)
	r.GET("/games/:id/observer.log", public.GetObserverLog)

	auth := r.Group("management", player.AuthRequired)
	auth.GET("/", player.GetPlayerSite)
	auth.POST("/update/", player.PostUpdate)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
