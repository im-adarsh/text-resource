package endpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/service"
)

const endpointPrefix = "/text-resource/v1"

const (
	Load    = endpointPrefix + "/load"
	GetText = endpointPrefix + "/get-text/:locale/:key"
)

func MakeEndPoint(r *gin.Engine, t service.TranslationService) {

	r.GET(Load, func(c *gin.Context) {
		err := t.Load()
		if err != nil {
			fmt.Println("", err)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET(GetText, func(c *gin.Context) {
		locale := c.Param("locale")
		key := c.Param("key")
		s, err := t.GetText(key, locale)
		if err != nil {
			fmt.Println("", err)
		}
		c.JSON(200, gin.H{
			"message": s,
		})
	})
}
