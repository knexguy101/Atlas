package nike

import (
	"strings"
	"time"
)

func (t *Task) login() int {

	t.setStatus("logging in")

	t.Delay()

	if !strings.Contains(t.page.URL(), "https://accounts.nike.com/lookup?") {
		err := t.page.Click("button[data-qa] > svg")
		if err != nil {
			return t.Error(err)
		}

		t.Delay()

		err = t.page.Click("button[data-qa='join-login-button']")
		if err != nil {
			return t.Error(err)
		}

		t.Delay()
	}

	t.page.Evaluate(`document.querySelector('#username').value = ''`)
	err := t.page.Type("#username", t.Account.Email)
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	err = t.page.Click("button[aria-label='Next']")
	if err != nil {
		return t.Error(err)
	}

	time.Sleep(35 * time.Second)

	err = t.page.Type("#password", t.Account.Password)
	if err != nil {
		return t.Error(err)
	}

	err = t.page.Click("button[aria-label='Sign In']")
	if err != nil {
		return t.Error(err)
	}

	time.Sleep(10 * time.Second) //fuck it

	c, err := t.page.Context().Cookies()
	if err != nil {
		return t.Error(err)
	}

	err = t.Account.Save(c)
	if err != nil {
		return t.Error(err)
	}

	if t.IsProfileTask {
		return ADDRESS
	} else {
		return CLEAR
	}
}
