package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()

	homeFP, err := os.Open("templates/home.plush.html")
	if err != nil {
		log.Fatalf("Error opening home template: %s", err)
	}
	defer homeFP.Close()
	homeTpl, err := ioutil.ReadAll(homeFP)
	if err != nil {
		log.Fatalf("Error reading home template: %s", err)
	}
	imgFP, err := os.Open("templates/image.plush.html")
	if err != nil {
		log.Fatalf("Error opening image template %s", err)
	}
	defer imgFP.Close()
	imgTpl, err := ioutil.ReadAll(imgFP)
	if err != nil {
		log.Fatalf("Error eading image template: %s", err)
	}
	defer imgFP.Close()
	router.Static("/static", "./static")
	router.GET("/", mainHandler(string(homeTpl)))
	router.GET("/cats", catHandler(string(imgTpl)))
	router.GET("/dogs", dogHandler(string(imgTpl)))
	router.Run(":8080")
}
