package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
	"fmt"
)

func Assign(user users.User, id int) (res []string, chatId []int64, err error) {
	taskId, errorID := task.GetTaskId(id)
	if errorID != nil {
		err = fmt.Errorf("нет такой задачи")
		return []string{}, []int64{}, err
	}

	if task.AllTasks[taskId].Assignee != "" || task.AllTasks[taskId].Creator != user.UserName {
		var userId int
		var errorUserID error

		if task.AllTasks[taskId].Assignee != "" {
			userId, errorUserID = users.GetUserId(task.AllTasks[taskId].Assignee)
		} else {
			userId, errorUserID = users.GetUserId(task.AllTasks[taskId].Creator)
		}

		if errorUserID != nil {
			return []string{}, []int64{}, errorUserID
		}

		users.AllUsers[userId].DeleteTask(task.AllTasks[taskId].Id)
		str := "Задача \"" + task.AllTasks[taskId].Name + "\" назначена на @" + user.UserName // сообщение новому владельцу задачи
		res = append(res, str)
		chatId = append(chatId, users.AllUsers[userId].ChatId)
	}
	task.AllTasks[taskId].Assignee = user.UserName

	userId, errorUserID := users.GetUserId(user.UserName)
	if errorUserID != nil {
		err = fmt.Errorf("")
		return []string{}, []int64{}, err
	}

	if !user.IsUserHasTask(task.AllTasks[taskId].Id) {
		users.AllUsers[userId].UserTasks = append(users.AllUsers[userId].UserTasks, task.AllTasks[taskId].Id)
	}

	str := "Задача \"" + task.AllTasks[taskId].Name + "\" назначена на вас" // сообщение новому владельцу задачи
	res = append(res, str)
	chatId = append(chatId, user.ChatId)

	return res, chatId, nil
}
