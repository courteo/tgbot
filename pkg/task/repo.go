package task

import (
	"fmt"
	"strconv"
)

var AllTasks []Task

func GetTaskId(id int) (int, error) {
	for i, task := range AllTasks {
		if task.Id == id {
			return i, nil
		}
	}
	err := fmt.Errorf("нет такой задачи")
	return -1, err
}

func IsTaskContain(taskName string) bool {
	for _, task := range AllTasks {
		if task.Name == taskName {
			return true
		}
	}
	return false
}

func PrintTaskWithAssignee(currTask Task) string {
	return strconv.Itoa(currTask.Id) + ". " + currTask.Name + " by @" + currTask.Creator + "\n" +
		"/unassign_" + strconv.Itoa(currTask.Id) + " /resolve_" + strconv.Itoa(currTask.Id)
}

func PrintTaskWithoutAssignee(currTask Task) string {
	return strconv.Itoa(currTask.Id) + ". " + currTask.Name + " by @" + currTask.Creator + "\n" +
		"/assign_" + strconv.Itoa(currTask.Id)
}
