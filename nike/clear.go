package nike

import "github.com/playwright-community/playwright-go"

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

	for err == nil {
		err = t.page.Click("button[data-automation='remove-item-button']", playwright.PageClickOptions{
			Timeout: playwright.Float(500),
		})

		t.Delay()
	}

	return ATC
}
