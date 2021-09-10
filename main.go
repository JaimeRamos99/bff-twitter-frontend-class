package main

import (
	structs "bff-twitter-frontend-class/tweets"
	http "net/http"
	time "time"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	tweets := structs.GetTweets()
	router.GET("/tweets", func(c *gin.Context) {
		c.JSON(http.StatusOK, tweets)
	})
	router.Run()
}
