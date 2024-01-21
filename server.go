package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startServer(config Config) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", pingHandler)

	router.GET("/", homeHandler)

	router.Run(fmt.Sprintf(":%d", config.Port))

}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "main website",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
