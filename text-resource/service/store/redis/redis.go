package redis

import (
	"strings"

	"github.com/go-redis/redis"
)

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(config Config) RedisClient {

	r := redisClient{}
	r.client = redis.NewClient(&redis.Options{
		Addr: "localhost:32768",
	})

	return &r
}

func (r *redisClient) GetTranslation(t GetTranslation) (string, error) {
	keyMap := map[string]string{
		Key:    t.Key,
		Locale: t.Locale,
	}

	key := GetKey(TranslationKey, keyMap)
	rs := r.client.Get(key)
	if rs.Err() != nil {
		return t.Key, rs.Err()
	} else {
		return rs.Val(), nil
	}
}

func (r *redisClient) SetTranslation(t SetTranslation) error {

	keyMap := map[string]string{
		Key:    t.Key,
		Locale: t.Locale,
	}

	key := GetKey(TranslationKey, keyMap)
	err := r.client.Set(key, t.Value, 0).Err()
	if err != nil {
		return err
	}

	if t.IsDefault {
		key := GetKey(TranslationKeyDefault, keyMap)
		err := r.client.Set(key, t.Value, 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *redisClient) Close() error {
	return r.client.Close()
}

func GetKey(rawKey string, valueMap map[string]string) string {
	if rawKey == "" || len(valueMap) < 1 {
		return rawKey
	}
	actualKey := rawKey
	for placeHolder, actualValue := range valueMap {
		actualKey = strings.Replace(actualKey, "{"+placeHolder+"}", actualValue, 1)
	}
	return actualKey
}
