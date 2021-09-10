package main

import (
	structs "bff-twitter-frontend-class/tweets"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	tweets := structs.GetTweets()
	router.GET("/tweets", func(c *gin.Context) {
		c.JSON(http.StatusOK, tweets)
	})
	router.Run()
}
