package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func (ts *TaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(taskStateJSON_marshal[*ts])
}

func (ts *TaskState) UnmarshalJSON(data []byte) error {
	if data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("improper TaskState %s in JSON", data)
	}
	data = data[1 : len(data)-1]
	_, ok := taskStateJSON_unmarshal[string(data)]
	if !ok {
		return fmt.Errorf("improper TaskState %s in JSON", data)
	}
	*ts = taskStateJSON_unmarshal[string(data)]

	return nil
}

func (tlist TaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlist.taskSlice)
}

func (tlist *TaskList) cleanUpHashLocation() error {
	tlist.hashLocation = make(map[TaskId]int)
	for pos, v := range tlist.taskSlice {
		_, bad := tlist.hashLocation[v.Id]
		if bad {
			return errors.New("provided JSON includes multiple tasks with the same TaskId")
		}
		tlist.hashLocation[v.Id] = pos
	}
	return nil
}

func (tlist *TaskList) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &tlist.taskSlice)
	if err != nil {
		return err
	}

	err = tlist.cleanUpHashLocation()
	if err != nil {
		fmt.Printf("Invalid JSON\n")
		log.Fatalf("Error trying to unmarshal json: %v\n", err)
	}
	return nil
}

func ReadFromFile(tl *TaskList, filename string) error {
	jsonFile, err := os.Open(filename)

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
	jsonFile, err := os.Create(filename)

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
