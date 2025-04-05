package main

import (
	"io"
	"log"

	"github.com/ursaru-tudor/task-cli/internal/cli"
)

func main() {
	log.SetOutput(io.Discard)
	app := cli.CreateApplication("task.json")
	app.ParseArguments()
}
