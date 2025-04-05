package task

import "encoding/json"

type TaskList struct {
	taskSlice    []Task
	hashLocation map[TaskId]int
}

func (tlist *TaskList) AddTask(t Task) {
	AssignTaskId(tlist.taskSlice, &t)
	tlist.taskSlice = append(tlist.taskSlice, t)

	if tlist.hashLocation == nil {
		tlist.hashLocation = make(map[TaskId]int)
	}

	tlist.hashLocation[t.Id] = len(tlist.taskSlice) - 1
}

func (tlist *TaskList) DeleteTask(id TaskId) {
	position := tlist.hashLocation[id]
	delete(tlist.hashLocation, id)
	tlist.taskSlice[position] = tlist.taskSlice[len(tlist.taskSlice)-1]
	tlist.taskSlice = tlist.taskSlice[:len(tlist.taskSlice)-1]
	tlist.hashLocation[tlist.taskSlice[position].Id] = position
}

func (tlist TaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlist.taskSlice)
}

func (tlist *TaskList) cleanUpHashLocation() {
	tlist.hashLocation = make(map[TaskId]int)
	for pos, v := range tlist.taskSlice {
		tlist.hashLocation[v.Id] = pos
	}
}

func (tlist *TaskList) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &tlist.taskSlice)
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

func (tlist TaskList) GetTasksByState(tsf TaskStateField) []TaskId {
	var r []TaskId
	for _, v := range tlist.taskSlice {
		if int(v.Status)&int(tsf) != 0 {
			r = append(r, v.Id)
		}
	}
	return r
}

func (tlist *TaskList) GetTask(id TaskId) *Task {
	index := tlist.hashLocation[id]
	return &tlist.taskSlice[index]
}
