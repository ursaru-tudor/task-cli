package task

import (
	"encoding/json"
	"sort"
	"time"
)

type TaskState int

const (
	TaskStateUnfinished TaskState = 1 << iota
	TaskStateActive     TaskState = 1 << iota
	TaskStateFinished   TaskState = 1 << iota
)

var taskStateName = map[TaskState]string{
	TaskStateUnfinished: "Unfinished",
	TaskStateActive:     "In Progress",
	TaskStateFinished:   "Finished",
}

func (ts TaskState) String() string {
	return taskStateName[ts]
}

type TaskId int

type Task struct {
	Id          TaskId
	Description string
	Status      TaskState
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) UpdateText(newText string) {
	t.Description = newText
	t.UpdatedAt = time.Now()
}

// Assigns a new unique Task Id
// Utilises mex (minimum excluded value)
// This is slow. The entire Task Id system is slow.
// But the number of tasks should be small enough for that not to matter.
func AssignTaskId(tslice []Task, t *Task) {
	mex := 1

	numbersUsed := []int{0}

	for _, v := range tslice {
		numbersUsed = append(numbersUsed, int(v.Id))
	}

	sort.Slice(numbersUsed, func(i, j int) bool {
		return numbersUsed[i] < numbersUsed[j]
	})

	for _, v := range numbersUsed {
		if v == mex {
			mex = v + 1
		}
	}

	t.Id = TaskId(mex)
}

func CreateTask(tDesc string) Task {
	var t Task
	t.Description = tDesc
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.Status = TaskStateUnfinished
	return t
}

type TaskList struct {
	TaskSlice    []Task
	hashLocation map[TaskId]int
}

func (tlist *TaskList) AddTask(t Task) {
	AssignTaskId(tlist.TaskSlice, &t)
	tlist.TaskSlice = append(tlist.TaskSlice, t)

	if tlist.hashLocation == nil {
		tlist.hashLocation = make(map[TaskId]int)
	}

	tlist.hashLocation[t.Id] = len(tlist.TaskSlice) - 1
}

func (tlist *TaskList) DeleteTask(id TaskId) {
	position := tlist.hashLocation[id]
	delete(tlist.hashLocation, id)
	tlist.TaskSlice[position] = tlist.TaskSlice[len(tlist.TaskSlice)-1]
	tlist.TaskSlice = tlist.TaskSlice[:len(tlist.TaskSlice)-1]
	tlist.hashLocation[tlist.TaskSlice[position].Id] = position
}

func (tlist TaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlist.TaskSlice)
}

func (tlist *TaskList) cleanUpHashLocation() {
	tlist.hashLocation = make(map[TaskId]int)
	for pos, v := range tlist.TaskSlice {
		tlist.hashLocation[v.Id] = pos
	}
}

func (tlist *TaskList) UnmarshalJSON(data []byte) error {

	err := json.Unmarshal(data, &tlist.TaskSlice)
	if err != nil {
		return err
	}
	tlist.cleanUpHashLocation()
	return nil
}

type TaskStateField int

func (ts *TaskStateField) AddTodo() {
	(*ts) |= TaskStateField(TaskStateUnfinished)
}

func (tlist TaskList) GetTasks(tsf TaskStateField) []TaskId {
	var r []TaskId
	for _, v := range tlist.TaskSlice {
		if int(v.Status)&int(tsf) != 0 {
			r = append(r, v.Id)
		}
	}
	return r
}
