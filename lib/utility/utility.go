package utility

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Config struct {
	APIKey string
}

func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetAPIKey(args []string) (string, error) {
	data, err := ioutil.ReadFile("./config.json")

	if err == nil {
		var conf Config

		err := json.Unmarshal(data, &conf)

		if err == nil && len(conf.APIKey) > 0 {
			return conf.APIKey, nil
		}

		return "", errors.New("Некорректный файл конфигурации")
	}

	return "", errors.New("Нет файла конфигурации")
}

func Rjust(initial string, char rune, size int) string {
	initialSize := len(initial)

	if size > initialSize {
		return initial + strings.Repeat(string(char), size-initialSize)
	}

	return initial
}
