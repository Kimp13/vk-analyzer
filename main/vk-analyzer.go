package main

import (
	"os"

	"lib/commands"
	"lib/logger"
	"lib/network"
)

func main() {
	argsSize := len(os.Args)
	network.DefaultURL.Scheme = "https"
	network.DefaultURL.Host = "api.vk.com"
	network.DefaultURL.Path = "/method/"

	network.DefaultQuery.Set("v", "5.126")

	logger.SetErrorPrefix(logger.BoldRed.Sprint("ОШИБКА: "))

	if argsSize > 1 {
		commands.All["help"] = commands.Command{
			commands.Help,
			commands.All["help"].Description,
		}

		if command, ok := commands.All[os.Args[1]]; ok {
			command.Exec(os.Args)
			return
		}

		logger.Errorf("Неизвестная команда: %s\n", logger.BoldCyan.Sprint(os.Args[1]))
	} else {
		logger.Errorln("Нет команды для выполнения!")
	}

	logger.Errorf(
		"Выполните команду %s, чтобы просмотреть доступные команды.\n",
		logger.BoldCyan.Sprintf("%s help", os.Args[0]),
	)
}
