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
		log.Printf("Couldn't opened %s to read\n", filename)
		return err
	}

	log.Printf("Successfully opened %s to read\n", filename)
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(bytes, &tl); err != nil {
		log.Printf("Failed unmarshalling %s. %v\n", filename, err)
		return err
	}

	return nil
}

func WriteToFile(tl task.TaskList, filename string) error {
	jsonFile, err := os.Create("task.json")

	if err != nil {
		log.Printf("Couldn't opened %s to write\n", filename)
		return err
	}

	log.Printf("Successfully opened %s to write\n", filename)
	defer jsonFile.Close()

	jsonBytes, err := json.MarshalIndent(tl, "", "  ")

	if err != nil {
		log.Printf("Failed marshalling %s\n", filename)
		return err
	}

	n, err := jsonFile.Write(jsonBytes)

	if err != nil || n != len(jsonBytes) {
		log.Printf("Failed writing to %s\n", filename)
		return err
	}

	return nil
}

func main() {
	//log.SetOutput(io.Discard)

	var savefile string = "test_task.json"
	var myTasks task.TaskList

	err := ReadFromFile(&myTasks, savefile)
	if err != nil {
		fmt.Printf("Loaded file from %s\n", savefile)
	}

	t := task.CreateTask("Clean around the house")
	myTasks.AddTask(t)

	jsonBytes, err := json.MarshalIndent(myTasks, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling tasks: %v", err)
	}
	fmt.Printf("%s\n", jsonBytes)

	err = WriteToFile(myTasks, "task.json")
	if err != nil {
		fmt.Printf("Failed to save data! Information may be lost. %v", err)
	}
}
