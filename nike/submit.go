package nike

import (
	"time"
)

func (t *Task) submit() int {
	t.setStatus("Submitting checkout modal")

	frameElement, err := t.page.QuerySelector(".checkout-modal > iframe")
	if err != nil {
		return t.Error(err)
	}

	frame, err := frameElement.ContentFrame()
	if err != nil {
		return t.Error(err)
	}

	t.Delay()

	e, err := frame.Locator(`button:has-text("Order")`)
	if err != nil {
		return t.Error(err)
	} else if vis, err := e.IsVisible(); err != nil || !vis {
		return t.Error(err)
	}
	e.Click()

	time.Sleep(10 * time.Second)

	//close the Your In Line popup
	t.page.Mouse().Click(0, 1)

	t.Delay()

	return PENDING
}
