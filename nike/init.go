package nike

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

func (t *Task) init() int {
	t.setStatus("Initializing task")

	var proxyUser string
	var proxyPassword string

	opts := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	}
	if len(t.Proxy) == 2 {
		proxyServer := fmt.Sprintf("http://%s:%s", t.Proxy[0], t.Proxy[1])

		if len(t.Proxy) > 2 {
			proxyUser = t.Proxy[2]
			proxyPassword = t.Proxy[3]
		} else {
			proxyUser = ""
			proxyPassword = ""
		}
		opts.Proxy = &playwright.BrowserTypeLaunchOptionsProxy{Server: &proxyServer, Username: &proxyUser, Password: &proxyPassword}
	}

	browser, err := t.PW.Firefox.Launch(opts)
	if err != nil {
		return t.Error(err)
	}
	t.browser = browser

	page, err := browser.NewPage()
	if err != nil {
		return t.Error(err)
	}
	t.page = page

	page.SetViewportSize(800, 800)

	return LOAD
}
