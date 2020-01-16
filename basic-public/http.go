package main

import "github.com/gin-gonic/gin"

func renderHTML(c *gin.Context, str string) {
	c.Header("Content-Type", "text/html")
	c.String(200, str)

}
