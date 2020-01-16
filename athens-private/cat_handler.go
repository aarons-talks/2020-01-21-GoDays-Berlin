package main

import (
	"github.com/gin-gonic/gin"
	"github.com/arschles/godays2020private"
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
	return godays2020private.ImagesHandler("/cats", tpl, cats)
}
