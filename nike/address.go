package nike

import (
	"github.com/playwright-community/playwright-go"
	"time"
)

/*
	Address

	Attempts to fill in the address information of the account with the given profile
 */

func (t *Task) address() int {
	t.setStatus("adding address")

	_, err := t.page.Goto("https://www.nike.com/member/settings")
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	_, err = t.page.Goto("https://www.nike.com/member/settings/delivery-addresses")
	if err != nil {
		return t.Error(err)
	}
	t.Delay()

	//Popup handling
	l, err := t.page.Locator("button:has-text('Continue')")
	if err != nil {
		return t.Error(err)
	} else if val, _ := l.IsVisible(); val {
		l.Click()
	}

	//remove old payment
	t.page.Click("button[data-testid='edit-button']")
	time.Sleep(2 * time.Second)
	t.page.Click("button[data-testid='open-delete-modal']")
	time.Sleep(750 * time.Millisecond)
	t.page.Click("button[data-testid='confirm-delete']")
	time.Sleep(750 * time.Millisecond)

	t.page.WaitForSelector("button[data-testid='add-address']")
	t.Delay()
	t.page.Click("button[data-testid='add-address']")
	time.Sleep(time.Millisecond * 500)

	t.page.Click(`#nameGiven`)
	t.page.Keyboard().Type(t.Profile.Shipping.FirstName, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#nameFamily`)
	t.page.Keyboard().Type(t.Profile.Shipping.LastName, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#line1`)
	t.page.Keyboard().Type(t.Profile.Shipping.Address1, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#line2`)
	t.page.Keyboard().Type(t.Profile.Shipping.Address2, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#locality`)
	t.page.Keyboard().Type(t.Profile.Shipping.City, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#code`)
	t.page.Keyboard().Type(t.Profile.Shipping.Zip, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	//US, CA
	t.page.Click(`#country`)
	t.page.SelectOption("#country", playwright.SelectOptionValues{
		Values: playwright.StringSlice(t.Profile.Shipping.Country),
	})
	time.Sleep(time.Millisecond * 300)

	//FL, NY, NC, etc... BC for canada etc
	t.page.Click(`#province`)
	t.page.SelectOption("#province", playwright.SelectOptionValues{
		Values: playwright.StringSlice(t.Profile.Shipping.Province),
	})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#phone`)
	time.Sleep(time.Millisecond * 300)
	t.page.Keyboard().Type(t.Profile.Shipping.Phone, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click("#preferred")
	time.Sleep(time.Millisecond * 774)

	t.page.Click("button[data-testid='submit-button']")
	t.Delay()

	//div You Entered
	//this handles the address suggestion, picks the original address
	e, err := t.page.Locator("div[data-testid='entered-address']")
	if err != nil {
		return t.Error(err)
	} else if val, _ := e.IsVisible(); val {
		e.Click()
		time.Sleep(time.Millisecond * 774)
	}

	t.page.Click("button[data-testid='submit-button']")
	t.Delay()

	return PAYMENT
}