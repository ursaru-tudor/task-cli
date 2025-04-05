package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ursaru-tudor/task-cli/internal/task"
)

func ReadFromFile(tl *task.TaskList, filename string) error {
	jsonFile, err := os.Open("task.json")

	if err != nil {
		log.Println("Couldn't opened task.json to read")
		return err
	}

	log.Println("Successfully opened task.json to read")
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(bytes, &tl); err != nil {
		log.Println("Failed unmarshalling task.json. Error: ", err)
		return err
	}

	return nil
}

func WriteToFile(tl task.TaskList, filename string) error {
	jsonFile, err := os.Create("task.json")

	if err != nil {
		log.Println("Couldn't opened task.json to write")
		return err
	}

	log.Println("Successfully opened task.json to write")
	defer jsonFile.Close()

	jsonBytes, err := json.MarshalIndent(tl, "", "  ")

	if err != nil {
		log.Println("Failed marshalling task.json")
		return err
	}

	n, err := jsonFile.Write(jsonBytes)

	if err != nil || n != len(jsonBytes) {
		log.Println("Failed writing to task.json")
		return err
	}

	return nil
}

func main() {
	println("This is the start of a good project :)")

	var myTasks task.TaskList

	ReadFromFile(&myTasks, "task.json")

	t := task.CreateTask("Clean around the house")
	myTasks.AddTask(t)

	jsonBytes, err := json.MarshalIndent(myTasks, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling tasks: %v", err)
	}
	fmt.Printf("%s\n", jsonBytes)

	WriteToFile(myTasks, "task.json")
}
