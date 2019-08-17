package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/service"
)

func main() {
	r := gin.Default()
	t := service.NewTranslationService()
	service.MakeEndPoint(r, t)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("unable to start text resource server")
	}
}
