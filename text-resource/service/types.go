package service

import (
	"context"

	"github.com/im-adarsh/text-resource/text-resource/text_resource_proto"
)

type TextResourceClient interface {
	GetText(ctx context.Context, req text_resource_proto.GetTextRequest) (text_resource_proto.GetTextResponse, error)
}
