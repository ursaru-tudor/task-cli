package cli

import "github.com/ursaru-tudor/task-cli/internal/task"

type Application struct {
	myTasks  task.TaskList
	savefile string
}

func CreateApplication(filename string) Application {
	var a Application
	a.savefile = filename
	task.ReadFromFile(&a.myTasks, a.savefile)
	return a
}

func (a Application) Save() {
	task.WriteToFile(a.myTasks, a.savefile)
}

// These functions assume already parsed input

func (a *Application) Add(text string) {
	t := task.CreateTask("Clean around the house")
	a.myTasks.AddTask(t)
}

func (a *Application) Update(id task.TaskId, text string) {
	taskHandler := a.myTasks.GetTask(id)
	taskHandler.UpdateText(text)
}

func (a *Application) Delete(id task.TaskId) {
	a.myTasks.DeleteTask(id)
}
