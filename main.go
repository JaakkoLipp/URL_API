package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type shorturl struct {
	ID			string `json:"id"`
	OriginalUrl string `json:"OriginalUrl"`
	ShortUrl	string `json:"ShortUrl"`
}

var urls = []shorturl{
	{ID: "1", OriginalUrl: "http://google.com", ShortUrl: "http://sh.fi/1"},
	{ID: "2", OriginalUrl: "http://bing.com", ShortUrl: "http://sh.fi/2"},
	{ID: "3", OriginalUrl: "https://duckduckgo.com", ShortUrl: "http://sh.fi/3"},
}

func main(){
	router := gin.Default()
	router.GET("/", root)
	router.GET("/urls", getUrls)

	router.Run("localhost:8080")
}

func getUrls(c *gin.Context){
	c.IndentedJSON(http.StatusOK, urls)
}

func root(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Usage: GET /urls for a json list"})
}
