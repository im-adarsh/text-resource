package redis

import "time"

const (
	Key    = "key"
	Locale = "locale"

	TranslationKey        = "translation_{key}_{locale}"
	TranslationKeyDefault = "translation_{key}_default"
)

type Config struct {
	Host    string
	Timeout time.Duration
}

type GetTranslation struct {
	Key    string `json:"key"`
	Locale string `json:"locale"`
}

type SetTranslation struct {
	Key       string `json:"key"`
	Locale    string `json:"locale"`
	Value     string `json:"value"`
	IsDefault bool   `json:"default"`
}

type RedisClient interface {
	GetTranslation(t GetTranslation) (string, error)
	SetTranslation(t SetTranslation) error
	Close() error
}
