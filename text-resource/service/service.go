package service

import "fmt"

type textResourceClient struct{}

func NewTextResourceClient()TextResourceClient {
	return &textResourceClient{}
}

func (textResourceClient) GetText(t GetTextRequest) string {
	res :=  t.textId
	// TODO get displayTextTmpl for cache or db
	res = fmt.Sprintf(t.textId, t.params...)
	return res

}

func (textResourceClient) AddText(t AddTextRequest) error {
	panic("implement me")
}
