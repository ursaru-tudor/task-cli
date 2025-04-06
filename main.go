package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/ursaru-tudor/task-cli/internal/cli"
)

const default_logging_file string = "task-cli.log"
const default_task_file string = "task.json"

//go:embed help.txt
var HelpString []byte

func printHelp() {
	fmt.Printf("%s\n", HelpString)
}

func main() {
	var logging_file string = default_logging_file
	var task_file string = default_task_file

	// All errors are saved in a log file
	logfile, err := os.Create(logging_file)
	if err != nil {
		log.Fatalf("Could not open logging file %s\n", logging_file)
	}
	log.SetOutput(logfile)

	if len(os.Args) < 2 || os.Args[1] == "help" {
		printHelp()
		return
	}

	app := cli.CreateApplication(task_file)

	app.ParseArguments()
}
