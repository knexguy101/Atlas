package nike

func (t *Task) pending() int {
	t.setStatus("Monitoring for win")

	_, err := t.page.Locator(`button:has-text("Pending")`)
	if err != nil {
		t.setStatus("Pending")
		t.Delay()
		return PENDING
	}

	_, err = t.page.Locator(`button:has-text("Buy")`)
	if err != nil {
		t.setStatus("Didn't get em")
		t.Delay()
		return FINISHED
	}

	t.setStatus("Check Email!")
	return FINISHED
}
