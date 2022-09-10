package account

import (
	"atlas/globals"
	"encoding/json"
	"sync"
)

type AccountList struct {
	sync.RWMutex
	List map[string]*Account
}

func NewAccountList() *AccountList {
	return &AccountList{
		List: make(map[string]*Account),
	}
}

func (al *AccountList) Add(acc *Account) {
	al.Lock()
	defer al.Unlock()

	al.List[acc.ID] = acc
	al.save()
}

func (al *AccountList) save() {
	b, err := json.Marshal(al.List)
	if err != nil {
		return
	}

	globals.SafeWriteFile(globals.AccountFilePath, b)
}

func (al *AccountList) Remove(id string) {
	al.Lock()
	defer al.Unlock()

	delete(al.List, id)
	al.save()
}

func (al *AccountList) Get(id string) (*Account, bool) {
	al.RLock()
	defer al.RUnlock()

	val, ok := al.List[id]
	return val, ok
}

func (al *AccountList) All() map[string]*Account {
	//lock not needed, maps are pointers, calling range on a map creates a separate object of the map.
	//same reason you can edit a map while looping through it with range.
	return al.List
}

func (al *AccountList) Load() {
	data, err := globals.SafeReadFile(globals.AccountFilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &al.List)
}