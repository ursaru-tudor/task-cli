package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

// Common error message

func ManageInvalidId(verb, argument string) {
	fmt.Printf("You have provided an invalid id (%s) to %s.\n", argument, verb)
	log.Printf("Error: Invalid TaskId %s for %s.\n", argument, verb)
}

// Parse arguments for each sub-command

func (a *Application) ParseList() {
	var tsf task.TaskStateField
	if len(os.Args) < 3 {
		tsf = task.TaskStateField(task.TaskStateActive | task.TaskStateFinished | task.TaskStateUnfinished)
	} else {
		for _, s := range os.Args[2:] {
			ls := strings.ToLower(s)
			switch ls {
			case "todo":
				tsf.AddState(task.TaskStateUnfinished)
			case "to-do":
				tsf.AddState(task.TaskStateUnfinished)
			case "inprogress":
				tsf.AddState(task.TaskStateActive)
			case "in-progress":
				tsf.AddState(task.TaskStateActive)
			case "done":
				tsf.AddState(task.TaskStateFinished)
			default:
				fmt.Printf("Invalid argument %s for list sumcommand\n", s)
				log.Printf("Invalid argument %s for list sumcommand\n", s)
				return
			}
		}
	}
	til := a.myTasks.GetTasksByState(tsf)
	if len(til) == 0 {
		fmt.Printf("No matching task found\n")
		return
	}

	fmt.Print(a.StringTasksLong(til))
}

func (a *Application) ParseAdd() {
	if len(os.Args) < 3 {
		fmt.Printf("Too few arguments for verb add. You must provide at least one title for a new task.\n")
		log.Printf("Error: Too few arguments for add.\n")
		return
	}

	for _, s := range os.Args[3:] {
		v := a.Add(s)
		fmt.Printf("Task added successfully. Id: %v\n", v)
	}

	a.Save()
}

func (a *Application) ParseInfo() {
	if len(os.Args) < 3 {
		fmt.Printf("Too few arguments for info. You must provide at least one TaskId.\n")
		log.Printf("Error: Too few arguments for info.\n")
		return
	}

	var tid []task.TaskId

	for _, s := range os.Args[2:] {
		id, err := task.ExtractIdFromString(s)
		if err != nil || !a.myTasks.CheckId(id) {
			ManageInvalidId("info", s)
			return
		}
		tid = append(tid, id)
	}

	for _, id := range tid {
		fmt.Println(TaskVerboseDisplay(*a.myTasks.GetTask(id)))
	}
}

func (a *Application) ParseDelete() {
	if len(os.Args) < 3 {
		fmt.Printf("Too few arguments for verb delete. You must provide at least one TaskId.\n")
		log.Printf("Error: Too few arguments for delete.\n")
		return
	}

	var tid []task.TaskId

	for _, s := range os.Args[2:] {
		id, err := task.ExtractIdFromString(s)
		if err != nil || !a.myTasks.CheckId(id) {
			ManageInvalidId("delete", s)
			return
		}
		tid = append(tid, id)
	}

	for _, id := range tid {
		a.Delete(id)
	}

	a.Save()
}

func (a *Application) ParseMark(ts task.TaskState) {
	if len(os.Args) < 3 {
		fmt.Printf("Too few arguments for verb class mark. You must provide at least one TaskId.\n")
		log.Printf("Error: Too few arguments for mark.\n")
		return
	}

	var tid []task.TaskId

	for _, s := range os.Args[2:] {
		id, err := task.ExtractIdFromString(s)
		if err != nil || !a.myTasks.CheckId(id) {
			ManageInvalidId("mark", s)
			return
		}
		tid = append(tid, id)
	}

	for _, id := range tid {
		a.Mark(id, ts)
	}

	a.Save()
}

func (a *Application) ParseArguments() {
	if len(os.Args) < 2 {
		fmt.Printf("You must include arguments to do anything. For information, check with 'task-cli help'\n")
		log.Printf("No arguments provided\n")
		return
	}
	verb := os.Args[1]
	verb = strings.ToLower(verb)
	switch verb {
	case "list":
		a.ParseList()
	case "add":
		a.ParseAdd()
	case "info":
		a.ParseInfo()
	case "delete":
		a.ParseDelete()
	case "mark-in-progress":
		a.ParseMark(task.TaskStateActive)
	case "mark-done":
		a.ParseMark(task.TaskStateFinished)
	default:
		fmt.Printf("You have included an invalid verb. For information on correct usage, check with 'task-cli help'\n")
		log.Printf("Invalid verb provided\n")
	}
}
