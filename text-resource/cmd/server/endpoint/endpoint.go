package endpoint

import (
	"context"
	"errors"
	"github.com/im-adarsh/text-resource/text-resource/text_resource_proto"

	"github.com/gin-gonic/gin"
	"github.com/im-adarsh/text-resource/text-resource/service"
)

const TextResourceUrlPrefix = "/text"

func RegisterEndPoint(r *gin.Engine, t service.TextResourceClient) {

	r.GET(TextResourceUrlPrefix+"/text", func(c *gin.Context) {
		t.GetText(service.GetTextRequest{})
		c.JSON(200, gin.H{
			"message": "pong get",
		})
	})

	r.POST(TextResourceUrlPrefix+"/text", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong post",
		})
	})

}

func MakeGetTextEndpoint(svc service.TextResourceClient) endpoin.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*text_resource_proto.GetTextRequest{})
		if !ok {
			return nil, errors.New("cant cast request to *cats_proto.AdRequest")
		}
		v, err := svc.GetText(ctx, req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
