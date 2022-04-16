package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func RenderError(c *gin.Context, where string, err error) {
	log.Printf("err @ %s %s %s: %s\n", c.Request.Method, c.FullPath(), where, err.Error())
	c.String(500, "err @ %s %s %s", c.Request.Method, c.FullPath(), where)
	c.Abort()
}

func Render(c *gin.Context, name string, data gin.H) {
	c.HTML(200, name, data)
}
