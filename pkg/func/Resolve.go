package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
	"fmt"
)

func Resolve(user users.User, id int) (res []string, chatId []int64, err error) {
	taskId, errorID := task.GetTaskId(id)
	if errorID != nil {
		err = fmt.Errorf("нет такой задачи")
		return []string{}, []int64{}, err
	}

	Assignee, errorUser := users.GetUserId(task.AllTasks[taskId].Assignee)
	if errorUser != nil {
		errorUser = fmt.Errorf("Нет пользователя, которому задали эту задачу")
		return []string{}, []int64{}, errorUser
	}

	if users.AllUsers[Assignee].UserName != user.UserName {
		err = fmt.Errorf("у вас нет доступка к этому")
		return []string{}, []int64{}, err
	}
	users.AllUsers[Assignee].DeleteTask(task.AllTasks[taskId].Id) // удаляем задачу у исполнителя
	str := "Задача \"" + task.AllTasks[taskId].Name + "\" выполнена"
	res = append(res, str)
	chatId = append(chatId, users.AllUsers[Assignee].ChatId)

	creator, errorUser := users.GetUserId(task.AllTasks[taskId].Creator)
	if errorUser != nil {
		return []string{}, []int64{}, errorUser
	}

	if creator == Assignee {
		return res, chatId, nil
	}
	users.AllUsers[creator].DeleteCreatedTask(task.AllTasks[taskId].Id) // удаляем задачу у создателя
	str = "Задача \"" + task.AllTasks[taskId].Name + "\" выполнена @" + users.AllUsers[Assignee].UserName
	res = append(res, str)
	chatId = append(chatId, users.AllUsers[creator].ChatId)

	task.AllTasks = append(task.AllTasks[:taskId], task.AllTasks[taskId+1:]...)
	return res, chatId, nil
}
