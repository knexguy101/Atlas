package nike

import (
	"sync"
)

type NikeTaskList struct {
	sync.RWMutex
	List map[string]*Task
}

func NewNikeTaskList() *NikeTaskList {
	return &NikeTaskList{
		List: make(map[string]*Task),
	}
}

func (al *NikeTaskList) Add(acc *Task) {
	al.Lock()
	defer al.Unlock()

	al.List[acc.ID] = acc
}

func (al *NikeTaskList) Remove(id string) {
	al.Lock()
	defer al.Unlock()

	delete(al.List, id)
}

func (al *NikeTaskList) Get(id string) (*Task, bool) {
	al.RLock()
	defer al.RUnlock()

	val, ok := al.List[id]
	return val, ok
}

func (al *NikeTaskList) All() map[string]*Task {
	//lock not needed, maps are pointers, calling range on a map creates a separate object of the map.
	//same reason you can edit a map while looping through it with range.
	return al.List
}