package profile

import (
	"atlas/globals"
	"encoding/json"
	"sync"
)

type ProfileList struct {
	sync.RWMutex
	List map[string]*Profile
}

func NewProfileList() *ProfileList {
	return &ProfileList{
		List: make(map[string]*Profile),
	}
}

func (al *ProfileList) Add(acc *Profile) {
	al.Lock()
	defer al.Unlock()

	al.List[acc.ID] = acc
	al.save()
}

func (al *ProfileList) save() {
	b, err := json.Marshal(al.List)
	if err != nil {
		return
	}

	globals.SafeWriteFile(globals.ProfileFilePath, b)
}

func (al *ProfileList) Remove(id string) {
	al.Lock()
	defer al.Unlock()

	delete(al.List, id)
	al.save()
}

func (al *ProfileList) Get(id string) (*Profile, bool) {
	al.RLock()
	defer al.RUnlock()

	val, ok := al.List[id]
	return val, ok
}

func (al *ProfileList) All() map[string]*Profile {
	//lock not needed, maps are pointers, calling range on a map creates a separate object of the map.
	//same reason you can edit a map while looping through it with range.
	return al.List
}

func (al *ProfileList) Load() {
	data, err := globals.SafeReadFile(globals.ProfileFilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &al.List)
}