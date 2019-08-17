package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/im-adarsh/text-resource/text-resource/service/store/redis"
)

type translationService struct {
	redis redis.RedisClient
}

func NewTranslationService() TranslationService {

	r := translationService{}
	r.redis = redis.NewRedisClient(redis.Config{})

	return &r
}

func (t *translationService) Load() error {

	jsonFile, err := os.Open("../resources/translation.json")
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var translations Translations
	err = json.Unmarshal(byteValue, &translations)
	if err != nil {
		return err
	}

	for _, v := range translations.Translations {
		err := t.redis.SetTranslation(v)
		if err != nil {
			fmt.Println("error : ", err)
		}
	}

	return nil
}
