package task

import (
	"slices"
	"strconv"
)

type TaskList struct {
	taskSlice    []Task
	hashLocation map[TaskId]int
}

func (tlist *TaskList) AddTask(t Task) TaskId {
	AssignTaskId(tlist.taskSlice, &t)
	tlist.taskSlice = append(tlist.taskSlice, t)

	if tlist.hashLocation == nil {
		tlist.hashLocation = make(map[TaskId]int)
	}

	tlist.hashLocation[t.Id] = len(tlist.taskSlice) - 1
	return t.Id
}

func (tlist *TaskList) DeleteTask(id TaskId) {
	position := tlist.hashLocation[id]
	delete(tlist.hashLocation, id)

	if position == len(tlist.taskSlice)-1 {
		tlist.taskSlice = tlist.taskSlice[:len(tlist.taskSlice)-1]
		return
	}

	tlist.taskSlice[position] = tlist.taskSlice[len(tlist.taskSlice)-1]
	tlist.taskSlice = tlist.taskSlice[:len(tlist.taskSlice)-1]
	tlist.hashLocation[tlist.taskSlice[position].Id] = position
}

type TaskStateField int

const AllTaskStates TaskStateField = TaskStateField(TaskStateActive | TaskStateFinished | TaskStateUnfinished)

func (ts *TaskStateField) AddState(t TaskState) {
	(*ts) |= TaskStateField(t)
}

func (tlist TaskList) Matches(id TaskId, tsf TaskStateField) bool {
	state := tlist.GetTask(id).Status
	return ((tsf) & TaskStateField(state)) != 0
}

func (tlist TaskList) GetTasksByState(tsf TaskStateField) []TaskId {
	var r []TaskId
	for _, v := range tlist.taskSlice {
		if int(v.Status)&int(tsf) != 0 {
			r = append(r, v.Id)
		}
	}
	slices.Sort(r)
	return r
}

func (tlist *TaskList) GetTask(id TaskId) *Task {
	index := tlist.hashLocation[id]
	return &tlist.taskSlice[index]
}

func (tlist TaskList) CheckId(id TaskId) bool {
	_, ok := tlist.hashLocation[id]
	return ok
}

func ExtractIdFromString(str string) (TaskId, error) {
	numId, err := strconv.Atoi(str)
	return TaskId(numId), err
}
