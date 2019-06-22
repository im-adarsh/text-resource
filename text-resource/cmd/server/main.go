package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/cmd/server/endpoint"
)

func main() {
	r := gin.Default()
	endpoint.RegisterEndPoint(r)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("unable to run server")
	}
}
