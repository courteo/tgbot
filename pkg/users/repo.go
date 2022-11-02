package users

import (
	"MYtgbot/pkg/task"
	"fmt"
	"strconv"
)

func PrintAllTasks(user User) (res string, err error) {
	if len(task.AllTasks) == 0 {
		err = fmt.Errorf("Нет задач")
		return "", err
	}

	for i, tasks := range task.AllTasks {
		str := strconv.Itoa(tasks.Id) + ". " + tasks.Name + " by @" + tasks.Creator + "\n"
		if tasks.Assignee != "" { // задачу кто-то взял
			if tasks.Assignee == user.UserName {
				str += "assignee: я\n"
				str += "/unassign_" + strconv.Itoa(tasks.Id) + " /resolve_" + strconv.Itoa(tasks.Id)
			} else {
				str += "assignee: @" + tasks.Assignee
			}

		} else { // задачу никто не взял
			str += "/assign_" + strconv.Itoa(tasks.Id)
		}
		res += str
		if i != len(task.AllTasks)-1 {
			res += "\n" + "\n"
		}
	}
	return res, nil
}

var AllUsers []User
var Inc int

func GetUserId(userName string) (int, error) {
	for i, user := range AllUsers {
		if user.UserName == userName {
			return i, nil
		}
	}
	err := fmt.Errorf("нет пользователя")
	return -1, err
}

func GetUser(userName string) (User, error) {
	for _, user := range AllUsers {
		if user.UserName == userName {
			return user, nil
		}
	}
	err := fmt.Errorf("нет пользователя")
	return User{}, err
}

func (user *User) AddNewTask(newTask task.Task) {
	user.CreatedTasks = append(user.CreatedTasks, newTask.Id)
}

func (user *User) DeleteTask(taskName int) {
	index := -1
	for i, task := range user.UserTasks {
		if task == taskName {
			index = i
			break
		}
	}
	if index != -1 {
		user.UserTasks = append(user.UserTasks[:index], user.UserTasks[index+1:]...)
	}
}

func (user *User) DeleteCreatedTask(taskName int) {
	index := -1
	for i, task := range user.CreatedTasks {
		if task == taskName {
			index = i
			break
		}
	}
	if index != -1 {
		user.CreatedTasks = append(user.CreatedTasks[:index], user.CreatedTasks[index+1:]...)
	}
}

func (user User) IsUserHasTask(taskName int) bool {
	for _, userTask := range user.UserTasks {
		if userTask == taskName {
			return true
		}
	}
	return false
}
