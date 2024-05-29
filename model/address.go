package model

type Address struct {
	ID             uint32 `json:"id"`
	Place          uint8  `json:"place"`
	Street         string `json:"street"`
	Height         string `json:"height"`
	Floor          uint16 `json:"floor"`
	Department     uint16 `json:"department"`
	Tower          uint16 `json:"tower"`
	BetweenStreets string `json:"betweenStreets"`
	Observations   string `json:"observations"`
	Shipment       bool   `json:"shipment"`
	ProvinceID     uint32 `json:"provinceID"`
}
