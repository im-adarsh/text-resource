package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const endpointPrefix = "/text-resource/v1"

func MakeEndPoint(r *gin.Engine, t TranslationService) {

	r.GET(endpointPrefix+"/load", func(c *gin.Context) {
		err := t.Load()
		if err != nil {
			fmt.Println("", err)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET(endpointPrefix+"/get-text/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		s, err := t.GetText("text.hello_world", locale)
		if err != nil {
			fmt.Println("", err)
		}
		c.JSON(200, gin.H{
			"message": s,
		})
	})
}
