package nike

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"time"
)

func (t *Task) payment() int {
	t.setStatus("adding payment")

	_, err := t.page.Goto("https://www.nike.com/member/settings/payment-methods")
	if err != nil {
		return t.Error(err)
	}
	t.Delay()

	//remove old payment
	t.page.Click("div[class='Default-payment'] > div > div > button")
	time.Sleep(2 * time.Second)
	t.page.Click("div[data-testid='dialog-actions'] > div > div > button:nth-child(2)")
	time.Sleep(750 * time.Millisecond)
	t.page.Click(".mex-address-delete-button")
	time.Sleep(750 * time.Millisecond)

	t.page.WaitForSelector(".mex-add-new-payment-container > div > button[type='submit']")
	t.page.Click(".mex-add-new-payment-container > div > button[type='submit']")
	t.Delay()

	paymentFrame := t.findFrame()
	if paymentFrame == nil {
		return FINISHED
	}

	paymentFrame.Click("#creditCardNumber")
	t.page.Keyboard().Type(t.Profile.Payment.CC, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 1373)

	paymentFrame.Click("#expirationDate")
	t.page.Keyboard().Type(fmt.Sprintf(fixMonth(t.Profile.Payment.Month) + t.Profile.Payment.Year[2:]), playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 1073)

	paymentFrame.Click("#cvNumber")
	t.page.Keyboard().Type(t.Profile.Payment.CVV, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 973)

	t.page.Click(".mex-new-card > form > button")
	time.Sleep(time.Millisecond * 500)

	t.page.Click(`#firstName`)
	t.page.Keyboard().Type(t.Profile.Billing.FirstName, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	t.page.Click(`#lastName`)
	t.page.Keyboard().Type(t.Profile.Billing.LastName, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 300)

	//US, CA
	t.page.Click(`#country`)
	t.page.SelectOption("#country", playwright.SelectOptionValues{
		Values: playwright.StringSlice(t.Profile.Billing.Country),
	})
	time.Sleep(time.Millisecond * 1300)

	//TX, BC
	t.page.Click(`#state`)
	t.page.SelectOption("#state", playwright.SelectOptionValues{
		Values: playwright.StringSlice(t.Profile.Billing.Province),
	})
	time.Sleep(time.Millisecond * 1300)

	t.page.Click(`#address1`)
	t.page.Keyboard().Type(t.Profile.Billing.Address1, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	time.Sleep(time.Millisecond * 1300)

	t.page.Click(`#address2`)
	t.page.Keyboard().Type(t.Profile.Billing.Address2, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	t.page.Click(`#address2`)
	time.Sleep(time.Millisecond * 1300)

	t.page.Click(`#city`)
	t.page.Keyboard().Type(t.Profile.Billing.City, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	t.page.Click(`#city`)
	time.Sleep(time.Millisecond * 1300)

	t.page.Click(`#postalCode`)
	t.page.Keyboard().Type(t.Profile.Billing.Zip, playwright.KeyboardTypeOptions{Delay: playwright.Float(147)})
	t.page.Click(`#postalCode`)
	time.Sleep(time.Millisecond * 1300)

	t.page.Click("div[aria-labelledby='dialog-billing-address'] > section > div:nth-child(2) > div > div > button")
	time.Sleep(time.Millisecond * 1300)

	t.page.Click("#preferred")
	time.Sleep(time.Millisecond * 774)

	t.page.Click("div[aria-labelledby='dialog-add-payment-method'] > section > div:nth-child(2) > div > div > button")
	t.Delay()

	return FINISHED
}

func (t *Task) findFrame() playwright.Frame {
	ele, err := t.page.QuerySelector("iframe")
	if err != nil {
		return nil
	}
	f, err := ele.ContentFrame()
	if err != nil {
		return nil
	}
	return f
}

func fixMonth(month string) string {
	if len(month) == 1 {
		return "0" + month
	} else {
		return month
	}
}