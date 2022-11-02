package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
	"strconv"
)

func NewTask(taskName string, creator users.User) (res string) {
	if taskName == "" {
		return "Название задачи не может быть пустой"
	}

	if task.IsTaskContain(taskName) {
		return "the \"" + taskName + "\" task already exists"
	}

	users.Inc++
	newTask := task.Task{
		Name:    taskName,
		Creator: creator.UserName,
		Id:      users.Inc,
	}
	creator.AddNewTask(newTask)
	task.AllTasks = append(task.AllTasks, newTask)

	index, err := users.GetUserId(creator.UserName)
	if err == nil {
		users.AllUsers = append(users.AllUsers[:index], users.AllUsers[index+1:]...)
	}
	users.AllUsers = append(users.AllUsers, creator)

	return "Задача \"" + taskName + "\" создана, id=" + strconv.Itoa(users.Inc)
}
