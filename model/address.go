package model

type Address struct {
	ID             uint32 `json:"id"`
	Place          places `json:"place"`
	Street         string `json:"street"`
	Height         string `json:"height"`
	Floor          uint16 `json:"floor"`
	Department     uint16 `json:"department"`
	BetweenStreets string `json:"betweenStreets"`
	Observations   string `json:"observations"`
	ProvinceID     uint32 `json:"provinceID"`
}

type places string

const (
	work = "work"
	home = "home"
)
