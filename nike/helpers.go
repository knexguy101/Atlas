package nike

func (t *Task) closePageAndBrowser() {
	if t.page != nil {
		t.page.Close()
	}
	if t.browser != nil {
		t.browser.Close()
	}
}