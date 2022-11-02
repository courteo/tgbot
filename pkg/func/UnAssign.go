package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
	"fmt"
)

func UnAssign(user users.User, id int) (res []string, chatId []int64, err error) {
	taskId, errorID := task.GetTaskId(id)
	if errorID != nil {
		err = fmt.Errorf("нет такой задачи")
		return []string{}, []int64{}, err
	}

	if !user.IsUserHasTask(task.AllTasks[taskId].Id) {
		res = append(res, "Задача не на вас")
		chatId = append(chatId, user.ChatId)
		return res, chatId, nil
	}

	task.AllTasks[taskId].Assignee = ""
	userId, errorUserID := users.GetUserId(user.UserName)
	if errorUserID != nil {
		return []string{}, []int64{}, errorUserID
	}

	users.AllUsers[userId].DeleteTask(task.AllTasks[taskId].Id)
	str := "Принято" // сняли задачу с пользователя
	res = append(res, str)
	chatId = append(chatId, users.AllUsers[userId].ChatId)

	userId, errorUserID = users.GetUserId(task.AllTasks[taskId].Creator)
	if errorUserID != nil {
		return []string{}, []int64{}, errorUserID
	}

	users.AllUsers[userId].DeleteTask(task.AllTasks[taskId].Id)
	str = "Задача \"" + task.AllTasks[taskId].Name + "\" осталась без исполнителя" // сообщение автору задачи

	res = append(res, str)
	chatId = append(chatId, users.AllUsers[userId].ChatId)
	return res, chatId, nil
}
