package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unicode"
	"unicode/utf8"

	"lib/logger"
	"lib/network"
	"lib/utility"
)

type Command struct {
	Exec        func([]string)
	Description string
}

func Help(args []string) {
	var (
		commandsLen  = len(All)
		names        = make([]string, commandsLen)
		descriptions = make([]string, commandsLen)
		maxLen       = -1
		i            = 0
	)

	logger.Infoln("Команды:\n")

	for key, value := range All {
		currLen := len(key)

		if currLen > maxLen {
			maxLen = currLen
		}

		names[i] = key
		descriptions[i] = value.Description

		i++
	}

	for i := range names {

		logger.Infof(
			"\t%s %s\n\n",
			logger.BoldCyan.Sprint(utility.Rjust(names[i], ' ', maxLen)),
			descriptions[i],
		)
	}
}

func Init(args []string) {
	_, err := utility.GetAPIKey(args)

	if err == nil {
		var resp string

		logger.Infoln("Конфигурация уже настроена. ")
		logger.Info("Вы точно хотите продолжить? (Y\\n): ")

		fmt.Scanln(&resp)

		if r, _ := utf8.DecodeRuneInString(resp); unicode.ToLower(r) == 'n' {
			return
		}
	}

	conf := utility.Config{}
	wrong := true

	for wrong {
		logger.Info("Введите ключ VK API: ")
		fmt.Scanln(&conf.APIKey)

		address := network.DefaultURL
		query := network.DefaultQuery

		query.Set("access_token", conf.APIKey)

		address.Path += "account.getAppPermissions"
		address.RawQuery = query.Encode()

		resp := network.PermissionResponse{}
		err := utility.GetJSON(address.String(), &resp)

		if err == nil && resp.Response > 0 {
			wrong = false
		} else {
			logger.Errorln("Некорректный ключ!")
		}
	}

	b, err := json.Marshal(conf)

	if err != nil {
		logger.Errorln("Ошибка сериализации JSON!")
		logger.Errorln("Разработчик офигел - скажите ему об этом.")
		return
	}

	err = ioutil.WriteFile("./config.json", b, 0777)

	if err != nil {
		logger.Errorln("Ошибка записи файла конфигурации!")
		return
	}

	logger.Infoln("Конфигурация завершена!")
}

func Anal(args []string) {
	size := len(args)

	if size > 2 {

	} else {

	}
}

var All = map[string]Command{
	"help": {
		func(s []string) {},
		"Вывод всех команд с их описанием",
	},
	"init": {
		Init,
		"Пошаговая настройка конфигурации",
	},
}
