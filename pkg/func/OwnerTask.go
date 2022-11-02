package _func

import (
	"MYtgbot/pkg/task"
	"MYtgbot/pkg/users"
)

func OwnerTask(user users.User) (res string) {
	if len(user.CreatedTasks) == 0 {
		return "вы не создали задачи"
	}

	for i, userTask := range user.CreatedTasks {
		taskId, err := task.GetTaskId(userTask)
		if err != nil {
			return "нет такой задачи"
		}

		if task.AllTasks[taskId].Assignee != "" {
			res += task.PrintTaskWithAssignee(task.AllTasks[taskId])
		} else {
			res += task.PrintTaskWithoutAssignee(task.AllTasks[taskId])
		}

		if i != (len(user.CreatedTasks) - 1) {
			res += "\n"
		}
	}
	return res
}
