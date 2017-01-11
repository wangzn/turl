package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangzn/turl"
)

var (
	salt string
	addr string
	port string
	t    *turl.TURL
	err  error
)

func init() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.StringVar(&salt, "hash salt", "atinyurlservice", "salt string for hashing method")
	flag.StringVar(&addr, "redis addr", "redis://127.0.0.1:6379", "redis address for store")

}

//GetURL defines the method of get url from a key.
func GetURL(c *gin.Context) {
	key := c.Param("key")
	if len(key) == 0 {
		c.String(http.StatusNotFound, "Invalid request key")
	} else {
		url, e := t.GetURL(key)
		if e != nil {
			c.String(http.StatusInternalServerError, e.Error())
		} else {
			c.Redirect(http.StatusFound, url)
		}
	}
}

//AddURL defines the method of adding a url to turl, and returns the key.
func AddURL(c *gin.Context) {
	url := c.PostForm("url")
	if len(url) <= 0 {
		c.String(http.StatusBadRequest, "Invalid url")
	} else {
		key, e := t.Set(url)
		if e != nil {
			c.String(http.StatusInternalServerError, e.Error())
		}
		c.String(http.StatusOK, key)
	}
}

func main() {
	flag.Parse()
	fmt.Println(salt, addr)
	t, err = turl.New(salt, addr)
	if err != nil {
		panic(err)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/:key", GetURL)
	r.POST("/new", AddURL)
	r.Run(port)
}
