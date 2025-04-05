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
