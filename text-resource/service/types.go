package service

type GetTextRequest struct {
	textId string
	locale string
	params []interface{}
}

type AddTextRequest struct {
	textId string
	locale string
	displayTextTmpl string
}

type TextResourceClient interface {
	GetText(GetTextRequest) string
	AddText(AddTextRequest) error
}