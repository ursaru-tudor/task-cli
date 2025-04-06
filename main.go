package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ursaru-tudor/task-cli/internal/cli"
)

const default_logging_file string = "task-cli.log"
const default_task_file string = "task.json"

func printHelp() {
	helpFile, err := os.Open("help.txt")
	if err != nil {
		log.Fatalf("Failed to acquire help text\n")
	}

	str, err := io.ReadAll(helpFile)

	if err != nil {
		log.Fatalf("Failed to acquire help text\n")
	}

	fmt.Printf("%s\n", str)
}

func main() {
	var logging_file string = default_logging_file
	var task_file string = default_task_file

	logfile, err := os.Create(logging_file)
	if err != nil {
		log.Fatalf("Could not open logging file %s\n", logging_file)
	}
	log.SetOutput(logfile)

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	if os.Args[1] == "help" {
		printHelp()
		return
	}

	app := cli.CreateApplication(task_file)

	app.ParseArguments()
}
