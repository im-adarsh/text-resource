package service

import (
	"context"

	"github.com/im-adarsh/text-resource/text-resource/text_resource_proto"
)

type textResourceClient struct{}

func NewTextResourceClient() text_resource_proto.TextResourceServer {
	return &textResourceClient{}
}

func (textResourceClient) GetText(context.Context, *text_resource_proto.GetTextRequest) (*text_resource_proto.GetTextResponse, error) {
	panic("implement me")
}
