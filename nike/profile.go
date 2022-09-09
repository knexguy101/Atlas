package nike

type Profile struct {
	Shipping ProfileInfo `json:"shipping"`
	Billing ProfileInfo `json:"billing"`
	Payment ProfilePayment `json:"payment"`
}

type ProfileInfo struct {
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Address1 string `json:"a1"`
	Address2 string `json:"a2"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Province string `json:"province"`
	Country string `json:"country"`
	Phone string `json:"phone"`
}

type ProfilePayment struct {
	CC string `json:"cc"`
	Month string `json:"month"`
	Year string `json:"year"`
	CVV string `json:"cvv"`
}