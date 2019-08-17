package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/service"
)

func main() {
	r := gin.Default()
	t := service.NewTranslationService()
	r.GET("/ping", func(c *gin.Context) {
		err := t.Load()
		if err != nil {
			fmt.Println("", err)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("unable to start text resource server")
	}
}
