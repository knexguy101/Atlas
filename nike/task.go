package nike

import (
	"context"
	"github.com/playwright-community/playwright-go"
	"log"
	"math/rand"
	"time"
)

type Task struct {
	ID string
	SKU string
	Size *Size
	Proxy []string
	Account *Account
	Profile *Profile
	PW *playwright.Playwright
	IsProfileTask bool
	StartTime int64
	MonitorRelease bool

	state int
	status string
	running bool

	browser playwright.Browser
	page playwright.Page
	ctx context.Context
	cancel func()
}

func (t *Task) setStatus(s string) {
	log.Println(s)
	t.status = s
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