package task

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadFromFile(tl *TaskList, filename string) error {
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

func WriteToFile(tl TaskList, filename string) error {
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
