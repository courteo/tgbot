package ForBot

import (
	_func "MYtgbot/pkg/func"
	"MYtgbot/pkg/users"
	"fmt"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func ForCommand(bot tgbotapi.BotAPI, currUser users.User, update tgbotapi.Update) {
	var msg, command, body string
	var taskId int
	var errorConv error
	index := strings.Index(update.Message.Text, " ")
	if index != -1 {
		command = update.Message.Text[1:index]
		body = update.Message.Text[index+1:]
	} else {

		command = update.Message.Text[1:]
		taskIdTemp := strings.Index(command, "_")

		if taskIdTemp != -1 {
			taskId, errorConv = strconv.Atoi(command[taskIdTemp+1:])
			command = command[:taskIdTemp]

			if errorConv != nil {
				_, err := bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Следует вводить номер задачи",
				))
				if err != nil {
					fmt.Println("ошибка при отправке")
					return
				}
				return
			}
		}
	}
	var err1 error
	switch command {
	case "new":
		msg = _func.NewTask(body, currUser)
		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msg,
		))

	case "my":
		msg = _func.MyTask(currUser)
		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msg,
		))
	case "owner":
		msg = _func.OwnerTask(currUser)
		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msg,
		))
	case "assign":
		BotSend(bot, currUser, taskId, update, "assign")
	case "unassign":
		BotSend(bot, currUser, taskId, update, "unassign")
	case "resolve":
		BotSend(bot, currUser, taskId, update, "resolve")
	case "tasks":
		msg1, err := users.PrintAllTasks(currUser)
		if err != nil {
			msg1 = "Нет задач"
		}

		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			msg1,
		))
	case "start":
		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Введите /help"))

	case "help":
		Help(bot, currUser, update)
	default:
		_, err1 = bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Команды не существует",
		))

	}
	if err1 != nil {
		fmt.Println("ошибка при отправке")
		return
	}
}
