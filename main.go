package main

import (
	"atlas/globals"
	"atlas/models/account"
	"atlas/models/profile"
	"atlas/models/size"
	"atlas/models/task"
	"atlas/nike"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"github.com/zserge/lorca"
	"strings"
)

var (
	accList = account.NewAccountList()
	profList = profile.NewProfileList()
	taskList = task.NewTaskList()
	nikeTaskList = nike.NewNikeTaskList()
	proxyList []string
)

func init() {
	godotenv.Load()
	playwright.Install()
}

func main() {

	globals.LoadFiles()
	accList.Load()
	profList.Load()
	taskList.Load()
	proxyList, _ = globals.SafeReadLines(globals.ProxiesFilePath)

	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}

	ui, _ := lorca.New("", "", 1460, 800)
	defer ui.Close()

	ui.Bind("appLoaded", func() {
		for _, v := range accList.All() {
			d, err := json.Marshal(v)
			if err != nil {
				continue
			}
			ui.Eval(fmt.Sprintf(`window.store.commit("addAccount", %s)`, string(d)))
		}
		for _, v := range proxyList {
			ui.Eval(fmt.Sprintf(`window.store.commit("addProxy", "%s")`, v))
		}
		for _, v := range profList.All() {
			d, err := json.Marshal(v)
			if err != nil {
				continue
			}
			ui.Eval(fmt.Sprintf(`window.store.commit("addProfile", %s)`, string(d)))
		}
		for _, v := range taskList.All() {
			d, err := json.Marshal(v)
			if err != nil {
				continue
			}
			v.Status = "Stopped"

			acc, ok := accList.Get(v.AccountID)
			if !ok {
				continue
			}
			pr, ok := profList.Get(v.ProfileID)
			if !ok {
				continue
			}
			s := &size.Size{
				BaseValue: v.RawSize,
				Kids: v.RawSizeKids,
			}

			nt := nike.Task{
				Task: v,
				Account: acc,
				Profile: pr,
				Size: s,
				Proxy: strings.Split(v.RawProxy, ":"),
				PW: pw,
				UI: ui,
			}

			nikeTaskList.Add(&nt)

			ui.Eval(fmt.Sprintf(`window.store.commit("addTask", %s)`, string(d)))
		}
	})

	ui.Bind("addAccounts", func(lines []string) []*account.Account {
		var accs []*account.Account
		for _, v := range lines {
			splits := strings.Split(v, ":")
			if len(splits) < 2 {
				continue
			}
			acc := &account.Account{
				Email: splits[0],
				Password: splits[1],
				ID: uuid.NewString(),
			}
			accList.Add(acc)
			accs = append(accs, acc)
		}
		return accs
	})
	ui.Bind("deleteAccount", func(id string) bool {
		accList.Remove(id)
		return true
	})

	ui.Bind("saveProxies", func(lines []string) {
		proxyList = lines
		globals.SafeWriteFile(globals.ProxiesFilePath, []byte(strings.Join(lines, "\n")))
	})

	ui.Bind("getProfile", func(id string) *profile.Profile {
		p, ok := profList.Get(id)
		if !ok {
			return nil
		}

		return p
	})
	ui.Bind("deleteProfile", func(id string) bool {
		_, ok := profList.Get(id)
		if !ok {
			return false
		}

		profList.Remove(id)
		return true
	})
	ui.Bind("saveProfile", func(p profile.Profile) *profile.Profile {
		current, ok := profList.Get(p.ID)
		if ok && p.Title != current.Title {
			p.ID = uuid.NewString()
		}

		profList.Add(&p)
		return &p
	})

	ui.Bind("createTask", func(t task.Task) *task.Task {

		t.ID = uuid.NewString()
		p, ok := accList.Get(t.AccountID)
		if !ok {
			return nil
		}
		t.AccountEmail = p.Email
		t.Status = "Stopped"
		taskList.Add(&t)

		pr, ok := profList.Get(t.ProfileID)
		if !ok {
			return nil
		}
		s := &size.Size{
			BaseValue: t.RawSize,
			Kids: t.RawSizeKids,
		}

		nt := nike.Task{
			Task: &t,
			Account: p,
			Profile: pr,
			Size: s,
			Proxy: strings.Split(t.RawProxy, ":"),
			PW: pw,
			UI: ui,
		}

		nikeTaskList.Add(&nt)
		return &t
	})
	ui.Bind("stopTask", func(id string) bool {
		t, ok := nikeTaskList.Get(id)
		if !ok {
			return ok
		} else if !t.IsRunning() {
			return false
		}

		t.Stop()
		return true
	})
	ui.Bind("deleteTask", func(id string) bool {
		_, ok := taskList.Get(id)
		if !ok {
			return false
		}

		nt, ok := nikeTaskList.Get(id)
		if ok {
			nt.Stop()
		}

		nikeTaskList.Remove(id)
		taskList.Remove(id)
		return true
	})
	ui.Bind("startTask", func(id string) bool {
		t, ok := nikeTaskList.Get(id)
		if !ok {
			return false
		} else if t.IsRunning() {
			return false
		}

		go t.Run()
		return true
	})

	ui.Load(fmt.Sprintf("http://127.0.0.1:5173"))

	// Wait for the browser window to be closed
	<-ui.Done()
}