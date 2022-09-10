package nike

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"time"
)

func (t *Task) cart() int {

	_, err := t.page.Goto(fmt.Sprintf("https://www.nike.com/launch/r/%s", t.SKU))
	if err != nil {
		return t.Error(err)
	}
	t.Delay()
	t.page.Reload()
	time.Sleep(10 * time.Second)

	if t.StartTime > 0 {
		un := time.Until(time.Unix(t.StartTime, 3))
		t.setStatus("Sleeping until release")
		select {
		case <- t.ctx.Done():
		case <- time.After(un):
		}
	} else if t.MonitorRelease {
		t.setStatus("Monitoring for release")
		for t.running {
			l, err := t.page.Locator(`button:has-text("Buy")`)
			if err != nil {
				t.setStatus("Error searching for Buy button")
				t.Delay()
				t.page.Reload()
				t.Delay()
			} else if val, err := l.IsVisible(); err != nil {
				t.setStatus("Error searching for Buy button visibility")
				t.Delay()
				t.page.Reload()
				t.Delay()
			} else if !val {
				t.setStatus("Buy button not found")
				time.Sleep(10 * time.Second)
				t.page.Reload()
				t.Delay()
			} else {
				break
			}
		}
		if !t.running {
			return FINISHED
		}
	}

	t.setStatus("Release found")

	sizeText, err := t.Size.GetSizeStringByRegion("US")
	if err != nil {
		return t.Error(err)
	}
	fmt.Println(sizeText)

	res, err := t.page.Locator(fmt.Sprintf(`button:has-text("M %s")`, sizeText))
	if err != nil {
		return t.Error(err)
	} else if vis, err := res.IsVisible(); err != nil || !vis {
		return t.Error(err)
	}
	res.Click()

	t.Delay()

	res, err = t.page.Locator(`button:has-text("Buy")`)
	if err != nil {
		return t.Error(err)
	} else if vis, err := res.IsVisible(); err != nil || !vis {
		return t.Error(err)
	}
	res.Click()

	t.Delay()

	t.setStatus("Waiting for checkout modal")
	_, err = t.page.WaitForSelector(".checkout-modal", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(25000),
	})

	return SUBMIT
}
