package main

import (
	"encoding/json"
	"io/ioutil"
	"logger"
	"os"
	"strings"

	"github.com/fatih/color"
)

type command struct {
	exec        func([]string)
	description string
}

func rjust(initial string, char rune, size int) string {
	initialSize := len(initial)

	if size > initialSize {
		return initial + strings.Repeat(string(char), size-initialSize)
	}

	return initial
}

type config struct {
	APIKey string
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

	getAPIKey(args)
}

func initialize(args []string) {

}

func getAPIKey(args []string) (string, error) {
	data, err := ioutil.ReadFile("./config.json")

	if err == nil {
		var conf config

		err := json.Unmarshal(data, &conf)

		if err == nil && len(conf.APIKey) > 0 {
			Log.Infoln(conf.APIKey)
			return conf.APIKey, nil
		}

		Log.Errorln("Некорректный файл конфигурации!")
	} else {
		Log.Errorln("Нет файла конфигурации!")
	}

	Log.Errorf(
		"Выполните команду %s, чтобы настроить конфигурацию.\n",
		boldCyan.Sprintf("%s init", args[0]),
	)

	return "", err
}

// Log выступает глобальным главным логгером для ошибок и обычного вывода
var Log *logger.Logger
var boldRed, boldCyan *color.Color
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
