package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

func (a *Application) ParseList() {
	// Will need to implement code for specific listings
	tsf := task.TaskStateField(task.TaskStateActive | task.TaskStateFinished | task.TaskStateUnfinished)
	fmt.Print(a.StringTasksLong(a.myTasks.GetTasksByState(tsf)))
}

func ManageArgumentCount(desiredArgNumber int, verb string, title bool) {
	if len(os.Args) > desiredArgNumber+1 {
		fmt.Printf("Too many arguments for verb %s.\n", verb)
		if title {
			fmt.Printf("If you want to input a title with spaces, make sure to insert \" before and after the text.")
		}
		fmt.Printf("\n")
		log.Fatalf("Error: Too many arguments for %s.\n", verb)
	}
	if len(os.Args) < desiredArgNumber+1 {
		fmt.Printf("Too few arguments for verb %s. You must provide the title of the task as a third argument.\n", verb)
		log.Fatalf("Error: Too few arguments for %s.\n", verb)
	}
}

func ManageInvalidId(verb, argument string) {
	fmt.Printf("You have provided an invalid id (%s) to %s.\n", argument, verb)
	log.Fatalf("Error: Invalid TaskId %s for %s.\n", argument, verb)
}

func (a *Application) ParseAdd() {
	ManageArgumentCount(2, "add", true)
	title := os.Args[2]
	a.Add(title)
	fmt.Println("Successfuly created new task.")
	task.WriteToFile(a.myTasks, a.savefile)
}

func (a *Application) ParseInfo() {
	ManageArgumentCount(2, "info", false)
	numId, err := strconv.Atoi(os.Args[2])
	id := task.TaskId(numId)

	if err != nil {
		fmt.Printf("Invalid TaskId format.\n")
		log.Fatalf("Invalid TaskId format\n")
	}

	if !a.myTasks.CheckId(id) {
		ManageInvalidId("info", os.Args[2])
	}

	fmt.Println(TaskVerboseDisplay(*a.myTasks.GetTask(id)))
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
	case "info":
		a.ParseInfo()
	default:
		fmt.Printf("You have included an invalid verb. For information on correct usage, check with the -h or --help flag\n")
		log.Fatalf("Invalid verb provided\n")
	}
}
