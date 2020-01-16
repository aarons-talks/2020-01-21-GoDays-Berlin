package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
)

func catHandler(tpl string) func(*gin.Context) {
	cats := []string{
		"/static/images/cats/1.jpg",
		"/static/images/cats/2.jpg",
		"/static/images/cats/3.jpg",
		"/static/images/cats/4.jpg",
		"/static/images/cats/5.jpg",
		"/static/images/cats/6.jpg",
	}
	return func(c *gin.Context) {
		ctx := plush.NewContext()
		ctx.Set("imgSrc", randStr(cats))
		ctx.Set("curURL", "/cats")
		str, err := plush.Render(tpl, ctx)
		if err != nil {
			c.String(500, "Error rendering template: %s", err)
			return
		}
		renderHTML(c, str)
	}
}
