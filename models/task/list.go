package task

import (
	"atlas/globals"
	"encoding/json"
	"sync"
)

type TaskList struct {
	sync.RWMutex
	List map[string]*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		List: make(map[string]*Task),
	}
}

func (al *TaskList) Add(acc *Task) {
	al.Lock()
	defer al.Unlock()

	al.List[acc.ID] = acc
	al.save()
}

func (al *TaskList) save() {
	b, err := json.Marshal(al.List)
	if err != nil {
		return
	}

	globals.SafeWriteFile(globals.TasksFilePath, b)
}

func (al *TaskList) Remove(id string) {
	al.Lock()
	defer al.Unlock()

	delete(al.List, id)
	al.save()
}

func (al *TaskList) Get(id string) (*Task, bool) {
	al.RLock()
	defer al.RUnlock()

	val, ok := al.List[id]
	return val, ok
}

func (al *TaskList) All() map[string]*Task {
	//lock not needed, maps are pointers, calling range on a map creates a separate object of the map.
	//same reason you can edit a map while looping through it with range.
	return al.List
}

func (al *TaskList) Load() {
	data, err := globals.SafeReadFile(globals.TasksFilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &al.List)
}