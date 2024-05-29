package model

type PhoneNumber struct {
	ID          uint32 `json:"id"`
	AreaCode    uint16 `json:"areaCode"`
	PhoneNumber string `json:"phoneNumber"`
	AccountID   uint32 `json:"accountID"`
}
