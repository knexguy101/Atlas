package account

import (
	"atlas/globals"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"path"
)

type Account struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Cookies []fileCookie `json:"cookies"`
}

type fileCookie struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Domain string `json:"domain"`
	Path string `json:"path"`
}

func Load(email string) (*Account, error) {
	filePath := path.Join("cookies", fmt.Sprintf("%x.txt", md5.Sum([]byte(email))))
	d, err := globals.SafeReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var a Account
	err = json.Unmarshal(d, &a)
	return &a, err
}

func (a *Account) ToNetworkCookies() []playwright.BrowserContextAddCookiesOptionsCookies {
	var cookies []playwright.BrowserContextAddCookiesOptionsCookies
	for _, v := range a.Cookies {
		cookies = append(cookies, playwright.BrowserContextAddCookiesOptionsCookies{
			Name: playwright.String(v.Name),
			Value: playwright.String(v.Value),
			Path: playwright.String(v.Path),
			Domain: playwright.String(v.Domain),
		})
	}
	return cookies
}

func (a *Account) Save(cookies []*playwright.BrowserContextCookiesResult) error {

	a.createFolder("cookies")

	filePath := path.Join("cookies", fmt.Sprintf("%x.txt", md5.Sum([]byte(a.Email))))
	_, err := globals.SafeCreateFile(filePath)
	if err != nil {
		return err
	}

	fa := Account{
		Email: a.Email,
		Password: a.Password,
	}
	for _, v := range cookies {
		fa.Cookies = append(fa.Cookies, fileCookie{
			Name: v.Name,
			Value: v.Value,
			Path: v.Path,
			Domain: v.Domain,
		})
	}

	d, err := json.Marshal(fa)
	if err != nil {
		return err
	}

	return globals.SafeWriteFile(filePath, d)
}

func (a *Account) createFolder(folder string) {
	globals.SafeCreateFolder(folder)
}