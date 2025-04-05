package cli

import (
	"fmt"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

func TaskShortDisplay(t task.Task) string {
	return fmt.Sprintf("%v) %v", t.Id, t.Description)
}

func TaskVerboseDisplay(t task.Task) string {
	fstr := "%v)\tName: %v" + "\n\tStatus: %v" + "\n\tCreated: %v" + "\n\tLast Updated: %v"
	return fmt.Sprintf(fstr, t.Id, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.UpdatedAt.Format("2006-01-02"))
}

func (a *Application) StringTasksShort(tsl []task.TaskId) string {
	var str string
	for _, v := range tsl {
		str += (TaskShortDisplay(*a.myTasks.GetTask(v)))
		str += "\n"
	}
	return str
}

func (a *Application) StringTasksLong(tsl []task.TaskId) string {
	var str string
	for _, v := range tsl {
		str += (TaskVerboseDisplay(*a.myTasks.GetTask(v)))
		str += "\n"
	}
	return str
}
