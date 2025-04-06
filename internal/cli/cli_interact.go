package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

// Common error messages

func ManageArgumentCount(desiredArgNumber int, verb string, title bool) bool {
	if len(os.Args) > desiredArgNumber+1 {
		fmt.Printf("Too many arguments for verb %s.\n", verb)
		if title {
			fmt.Printf("If you want to input a title with spaces, make sure to insert \" before and after the text.")
		}
		fmt.Printf("\n")
		log.Printf("Error: Too many arguments for %s.\n", verb)
		return false
	}
	if len(os.Args) < desiredArgNumber+1 {
		fmt.Printf("Too few arguments for verb %s. You must provide the title of the task as a third argument.\n", verb)
		log.Printf("Error: Too few arguments for %s.\n", verb)
		return false
	}
	return true
}

func ManageInvalidId(verb, argument string) {
	fmt.Printf("You have provided an invalid id (%s) to %s.\n", argument, verb)
	log.Printf("Error: Invalid TaskId %s for %s.\n", argument, verb)
}

// Parse arguments for each sub-command

// This entire thing needs to be rewritten
func (a *Application) ParseList() {
	// Will need to implement code for specific listings
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
	if !ManageArgumentCount(2, "add", true) {
		return
	}
	title := os.Args[2]
	v := a.Add(title)
	fmt.Printf("Task added successfully. Id: %v\n", v)
	a.Save()
}

func (a *Application) ParseInfo() {
	if !ManageArgumentCount(2, "info", false) {
		return
	}
	id, err := task.ExtractIdFromString(os.Args[2])

	if err != nil || !a.myTasks.CheckId(id) {
		ManageInvalidId("info", os.Args[2])
		return
	}

	fmt.Println(TaskVerboseDisplay(*a.myTasks.GetTask(id)))
}

func (a *Application) ParseDelete() {
	if !ManageArgumentCount(2, "delete", false) {
		return
	}
	id, err := task.ExtractIdFromString(os.Args[2])

	if err != nil || !a.myTasks.CheckId(id) {
		ManageInvalidId("delete", os.Args[2])
		return
	}

	a.Delete(id)
	a.Save()
}

func (a *Application) ParseMark(ts task.TaskState) {
	if !ManageArgumentCount(2, "mark", false) {
		return
	}

	id, err := task.ExtractIdFromString(os.Args[2])

	if err != nil || !a.myTasks.CheckId(id) {
		ManageInvalidId("mark", os.Args[2])
		return
	}

	a.Mark(id, ts)
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
		fmt.Printf("You have included an invalid verb. For information on correct usage, check with the -h or --help flag\n")
		log.Printf("Invalid verb provided\n")
	}
}
