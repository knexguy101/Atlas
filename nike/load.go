package nike

import (
	"atlas/models/account"
)

func (t *Task) load() int {

	t.setStatus("loading in")

	var (
		tempAcc = &account.Account{}
		err error
	)
	tempAcc, err = account.Load(t.Account.Email)
	if err == nil && len(tempAcc.Cookies) > 0 {

		t.setStatus("loading account")

		t.Account = tempAcc

		for _, v := range t.Account.ToNetworkCookies() {
			t.page.Context().AddCookies(v)
		}

		t.Delay()

		_, err = t.page.Goto("https://www.nike.com/launch")
		if err != nil {
			return t.Error(err)
		}
		t.Delay()
		t.page.Reload()
		t.Delay()

		if _, err := t.page.Locator("small[aria-label*='Options for']"); err != nil {
			return LOGIN
		}

		if t.IsProfileTask {
			return ADDRESS
		} else {
			return CLEAR
		}
	}

	_, err = t.page.Goto("https://www.nike.com/")
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	_, err = t.page.Goto("https://www.nike.com/launch")
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	return LOGIN
}
