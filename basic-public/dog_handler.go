package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
)

func dogHandler(tpl string) func(*gin.Context) {
	dogs := []string{
		"/static/images/dogs/1.jpg",
		"/static/images/dogs/2.jpg",
		"/static/images/dogs/3.png",
		"/static/images/dogs/3.png",
		"/static/images/dogs/4.png",
		"/static/images/dogs/5.png",
		"/static/images/dogs/6.jpg",
		"/static/images/dogs/7.jpg",
	}
	return func(c *gin.Context) {
		randomDog := randStr(dogs)
		ctx := plush.NewContext()
		ctx.Set("imgSrc", randomDog)
		ctx.Set("curURL", "/dogs")
		rendered, err := plush.Render(tpl, ctx)
		if err != nil {
			c.String(500, "Error rendering template: %s", err)
			return
		}
		renderHTML(c, rendered)
	}
}
