package account

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
	"os"
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
	filePath := path.Join("accounts", fmt.Sprintf("%x.txt", md5.Sum([]byte(email))))
	d, err := os.ReadFile(filePath)
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

	a.createFolder("accounts")

	filePath := path.Join("accounts", fmt.Sprintf("%x.txt", md5.Sum([]byte(a.Email))))
	if _, err := os.Stat(filePath); err != nil {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		f.Close()
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

	return os.WriteFile(filePath, d, 0777)
}

func (a *Account) createFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		mer := os.Mkdir(folder, os.ModePerm)
		if mer != nil {
			log.Println(mer.Error())
		}
	}
}