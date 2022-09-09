package main

import (
	"atlas/nike"
	"github.com/playwright-community/playwright-go"
)

func init() {
	playwright.Install()
}

func main() {

	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}

	t := nike.Task{
		ID: "test",
		SKU: "DX2938-200",
		Size: &nike.Size{
			BaseValue: 10,
			Kids: false,
		},
		Proxy: []string{},
		Account: &nike.Account{
			Email: "chris.falb@hotmail.com",
			Password: "Roxyisfast123",
		},
		PW: pw,
	}
	t.Run()
}