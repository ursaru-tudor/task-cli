package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

func (a *Application) ParseList() {
	// Will need to implement code for specific listings
	tsf := task.TaskStateField(task.TaskStateActive | task.TaskStateFinished | task.TaskStateUnfinished)
	fmt.Print(a.StringTasksLong(a.myTasks.GetTasksByState(tsf)))
}

func (a *Application) ParseAdd() {
	// Will need to implement code
	if len(os.Args) > 3 {
		log.Fatalf("Too many arguments for verb add")
		fmt.Printf("Error: Too many arguments")
	}
	title := os.Args[2]
	a.myTasks.AddTask(task.CreateTask(title))
	fmt.Println("Successfuly created new task.")
	task.WriteToFile(a.myTasks, a.savefile)
}

func (a *Application) ParseArguments() {
	if len(os.Args) < 2 {
		fmt.Printf("You must include arguments to do anything. For information, utilise the -h or --help flag\n")
		log.Fatalf("No arguments provided\n")
	}
	verb := os.Args[1]
	verb = strings.ToLower(verb)
	switch verb {
	case "list":
		a.ParseList()
	case "add":
		a.ParseAdd()
	default:
		fmt.Printf("You have included an invalid verb. For information on correct usage, check with the -h or --help flag\n")
		log.Fatalf("Invalid verb provided\n")
	}
}
