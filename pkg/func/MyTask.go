package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
)

func MyTask(user users.User) (res string) {
	for i, userTask := range user.UserTasks {
		tasks, err := task.GetTaskId(userTask)
		if err != nil {
			return "нет такой задачи"
		}

		res += task.PrintTaskWithAssignee(task.AllTasks[tasks])
		if i != len(user.UserTasks)-1 {
			res += "\n"
		}
	}
	if len(user.UserTasks) == 0 {
		return "на вас нет задач"
	}
	return res
}
