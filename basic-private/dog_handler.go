package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/plush"
	"github.com/arschles/godays2020private"

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
	return godays2020private.ImagesHandler("/dogs", tpl, dogs)
}
