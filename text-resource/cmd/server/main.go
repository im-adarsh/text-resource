package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/cmd/server/endpoint"
	"github.com/im-adarsh/text-resource/text-resource/service"
)

func main() {

	// init
	trClient := service.NewTextResourceClient()
	// create gin server
	r := gin.Default()
	endpoint.MakeGetTextEndpoint(trClient)
	endpoint.RegisterEndPoint(r, trClient)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("unable to run server")
	}
	fmt.Println("hello my name is adarsh kumar")

}

func helloworld() {

}

func helloworld1() {
	fmt.Println("i am working in carousell")
	fmt.Println("this is my new key board")
	fmt.Println("i am liking this new keyboar")

	fmt.Println("i am working in carousell")
	fmt.Println("this is my exiting keyboard")
	fmt.Println("i am liking this exiting keyboard")
}
