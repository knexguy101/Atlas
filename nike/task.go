package nike

import (
	"atlas/models/account"
	"atlas/models/profile"
	"atlas/models/size"
	"atlas/models/task"
	"context"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/zserge/lorca"
	"log"
	"math/rand"
	"time"
)

type Task struct {
	*task.Task
	Account *account.Account
	Profile *profile.Profile
	Size *size.Size
	Proxy []string

	state int
	running bool

	UI lorca.UI
	PW *playwright.Playwright
	browser playwright.Browser
	page playwright.Page
	ctx context.Context
	cancel func()
}

func (t *Task) setStatus(s string) {
	fmt.Println(t.ID, s)
	t.UI.Eval(fmt.Sprintf(`window.store.commit("setTaskStatus", {
		"id": "%s",
		"msg": "%s"
	})`, t.ID, s))
	t.Status = s
}

func (t *Task) Start() {
	t.running = true
}

func (t *Task) Stop() {
	t.running = false
}

func (t *Task) IsRunning() bool {
	return t.running
}

func (t *Task) Error(msg interface{}) int {
	log.Println("[", t.ID, "]", "[ERROR]", msg)
	time.Sleep(4 * time.Second)
	return t.state
}

func (t *Task) Delay() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(4) + 4) * time.Second)
}