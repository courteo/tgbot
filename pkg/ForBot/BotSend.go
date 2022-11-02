package ForBot

import (
	_func "MYtgbot/pkg/func"
	"MYtgbot/pkg/users"
	"fmt"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
)

func BotSend(bot tgbotapi.BotAPI, currUser users.User, taskId int, update tgbotapi.Update, name string) {
	var msgs []string
	var chatId []int64
	var err error
	switch name {
	case "assign":
		msgs, chatId, err = _func.Assign(currUser, taskId)
	case "unassign":
		msgs, chatId, err = _func.UnAssign(currUser, taskId)
	case "resolve":
		msgs, chatId, err = _func.Resolve(currUser, taskId)
	}

	if err != nil {
		_, err2 := bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"нет такой задачи",
		))
		if err2 != nil {
			fmt.Println("ошибка при отправке")
			return
		}
		return
	}

	for i := range msgs {
		_, err1 := bot.Send(tgbotapi.NewMessage(
			chatId[i],
			msgs[i],
		))
		if err1 != nil {
			fmt.Println("ошибка при отправке")
			return
		}
	}
}

func Help(bot tgbotapi.BotAPI, currUser users.User, update tgbotapi.Update) {
	str := "Существующие команды:\n \t /tasks - выводит текущие задачи\n \t /new XXX - вы создаете новую задачу\n" +
		"\t /assign_$ID  - назначаете пользователя исполнителем задачи\n \t /unassign_$ID - снимаете задачу с текущего пользователя\n" +
		"\t /resolve_$ID - выполняется задача\n \t /my - выводит задачи, которые назначили на меня\n \t /owner - показывает задачи, созданные мной"

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		str,
	))
	if err != nil {
		fmt.Println("ошибка при отправке")
		return
	}
}
