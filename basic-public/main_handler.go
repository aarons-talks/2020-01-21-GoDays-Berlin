package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
)

func mainHandler(tpl string) func(*gin.Context) {
	return func(c *gin.Context) {
		plushCtx := plush.NewContext()
		rendered, err := plush.Render(tpl, plushCtx)
		if err != nil {
			c.String(500, "Couldn't render template: %s", err)
			return
		}
		renderHTML(c, rendered)
	}
}
