package model

type Card struct {
	ID           uint32 `json:"ID"`
	Number       string `json:"Number"`
	SecurityCode string `json:"SecurityCode"`
	DueDate      string `json:"DueDate"`
	NameOwner    string `json:"NameOwner"`
}
