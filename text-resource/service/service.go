package service

type textResourceClient struct{}

func NewTextResourceClient()TextResourceClient {
	return &textResourceClient{}
}

func (textResourceClient) GetText(GetTextRequest) string {
	panic("implement me")
}

func (textResourceClient) AddText(AddTextRequest) error {
	panic("implement me")
}
