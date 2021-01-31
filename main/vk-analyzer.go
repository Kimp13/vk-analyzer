package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"local/logger"

	"github.com/fatih/color"
)

type command struct {
	exec        func([]string)
	description string
}

type config struct {
	APIKey string
}

type permissionResponse struct {
	Response uint32
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func rjust(initial string, char rune, size int) string {
	initialSize := len(initial)

	if size > initialSize {
		return initial + strings.Repeat(string(char), size-initialSize)
	}

	return initial
}

func help(args []string) {
	var (
		commandsLen  = len(commands)
		names        = make([]string, commandsLen)
		descriptions = make([]string, commandsLen)
		maxLen       = -1
		i            = 0
	)

	Log.Infoln("Команды:\n")

	for key, value := range commands {
		currLen := len(key)

		if currLen > maxLen {
			maxLen = currLen
		}

		names[i] = key
		descriptions[i] = value.description

		i++
	}

	for i := range names {

		Log.Infof(
			"\t%s %s\n\n",
			boldCyan.Sprint(rjust(names[i], ' ', maxLen)),
			descriptions[i],
		)
	}
}

func initialize(args []string) {
	_, err := getAPIKey(args)

	if err == nil {
		var resp string

		Log.Infoln("Конфигурация уже настроена. ")
		Log.Info("Вы точно хотите продолжить? (Y\\n): ")

		fmt.Scanln(&resp)

		if r, _ := utf8.DecodeRuneInString(resp); unicode.ToLower(r) == 'n' {
			return
		}
	}

	conf := config{}
	wrong := true

	for wrong {
		Log.Info("Введите ключ VK API: ")
		fmt.Scanln(&conf.APIKey)

		address := defaultURL
		query := defaultQuery

		query.Set("access_token", conf.APIKey)

		address.Path += "account.getAppPermissions"
		address.RawQuery = query.Encode()

		resp := permissionResponse{}
		err := getJSON(address.String(), &resp)

		if err == nil && resp.Response > 0 {
			wrong = false
		} else {
			Log.Errorln("Некорректный ключ!")
		}
	}

	b, err := json.Marshal(conf)

	if err != nil {
		Log.Errorln("Ошибка сериализации JSON!")
		Log.Errorln("Разработчик офигел - скажите ему об этом.")
		return
	}

	err = ioutil.WriteFile("./config.json", b, 0777)

	if err != nil {
		Log.Errorln("Ошибка записи файла конфигурации!")
		return
	}

	Log.Infoln("Конфигурация завершена!")
}

func getAPIKey(args []string) (string, error) {
	data, err := ioutil.ReadFile("./config.json")

	if err == nil {
		var conf config

		err := json.Unmarshal(data, &conf)

		if err == nil && len(conf.APIKey) > 0 {
			return conf.APIKey, nil
		}

		return "", errors.New("Некорректный файл конфигурации")
	}

	return "", errors.New("Нет файла конфигурации")
}

// Log выступает глобальным главным логгером для ошибок и обычного вывода
var Log *logger.Logger
var boldRed, boldCyan *color.Color
var defaultURL = url.URL{}
var defaultQuery = defaultURL.Query()
var commands = map[string]command{
	"help": {
		func(s []string) {},
		"Вывод всех команд с их описанием",
	},
	"init": {
		initialize,
		"Пошаговая настройка конфигурации",
	},
}

func main() {
	argsSize := len(os.Args)
	boldRed = color.New(color.FgRed, color.Bold)
	boldCyan = color.New(color.FgCyan, color.Bold)
	defaultURL.Scheme = "https"
	defaultURL.Host = "api.vk.com"
	defaultURL.Path = "/method/"

	defaultQuery.Set("v", "5.126")

	Log = logger.New()
	Log.SetErrorPrefix(boldRed.Sprint("ОШИБКА: "))

	if argsSize > 1 {
		commands["help"] = command{
			help,
			commands["help"].description,
		}

		if command, ok := commands[os.Args[1]]; ok {
			command.exec(os.Args)
			return
		}

		Log.Errorf("Неизвестная команда: %s\n", boldCyan.Sprint(os.Args[1]))
	} else {
		Log.Errorln("Нет команды для выполнения!")
	}

	Log.Errorf(
		"Выполните команду %s, чтобы просмотреть доступные команды.\n",
		boldCyan.Sprintf("%s help", os.Args[0]),
	)
}
