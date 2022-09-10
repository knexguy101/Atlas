package task

type Task struct {
	ID string `json:"id"`
	AccountEmail string `json:"accountEmail"`
	SKU string `json:"sku"`
	RawSize float64 `json:"rawSize"`
	RawSizeKids bool `json:"rawSizeKids"`
	RawProxy string `json:"rawProxy"`
	AccountID string `json:"accountId"`
	ProfileID string `json:"profileId"`
	IsProfileTask bool `json:"isProfileTask"`
	StartTime int64 `json:"startTime"`
	MonitorRelease bool `json:"monitorRelease"`
	Status string `json:"status"`
}