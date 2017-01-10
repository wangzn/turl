package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
}

func GetUrl(c *gin.Context) {
}

func AddUrl(c *gin.Context) {
}

func main() {
	flag.Parse()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/:key", GetUrl)
	r.POST("/new", AddUrl)
	r.Run(port)
}
