package task

import (
	"sort"
	"time"
)

type TaskState int

const (
	TaskStateUnfinished TaskState = 1 << iota
	TaskStateActive     TaskState = 1 << iota
	TaskStateFinished   TaskState = 1 << iota
)

var taskStateJSON_marshal = map[TaskState]string{
	TaskStateUnfinished: "todo",
	TaskStateActive:     "in-progress",
	TaskStateFinished:   "done",
}

var taskStateJSON_unmarshal = map[string]TaskState{
	"todo":        TaskStateUnfinished,
	"in-progress": TaskStateActive,
	"done":        TaskStateFinished,
}

var taskStateName = map[TaskState]string{
	TaskStateUnfinished: "unfinished",
	TaskStateActive:     "in progress",
	TaskStateFinished:   "finished",
}

func (ts TaskState) String() string {
	return taskStateName[ts]
}

type TaskId int

type Task struct {
	Id          TaskId    `json:"id"`
	Description string    `json:"description"`
	Status      TaskState `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
