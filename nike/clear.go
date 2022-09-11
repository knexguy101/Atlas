package nike

import "github.com/playwright-community/playwright-go"

/*
	Clear

	Clears the cart
 */

func (t *Task) clear() int {

	t.setStatus("clearing cart")

	_, err := t.page.Goto("https://www.nike.com")
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	_, err = t.page.Goto("https://www.nike.com/cart")
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	//messy, could replace by looping through selector list instead
	for err == nil && t.running {
		err = t.page.Click("button[data-automation='remove-item-button']", playwright.PageClickOptions{
			Timeout: playwright.Float(1500),
		})

		t.Delay()
	}

	return ATC
}
