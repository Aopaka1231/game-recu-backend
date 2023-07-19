package main

import (
	"game-rec-back/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//InitDBを実行
	db.InitDB()

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello, world",
		})
	})

	r.Run(":8000")
}