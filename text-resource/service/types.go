package service

import "github.com/im-adarsh/text-resource/text-resource/service/store/redis"

type Translations struct {
	Translations []redis.SetTranslation `json:"translations"`
}

type TranslationService interface {
	Load() error
}
