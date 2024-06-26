package types

type place string
type role string

const (
	user   role  = "user"
	admin  role  = "admin"
	worker role  = "worker"
	work   place = "work"
	home   place = "house"
)

type Address struct {
	ID             uint32 `json:"id"`
	Place          place  `json:"place"`
	Street         string `json:"street"`
	Height         string `json:"height"`
	Floor          uint16 `json:"floor"`
	Department     uint16 `json:"department"`
	BetweenStreets string `json:"betweenStreets"`
	Observations   string `json:"observations"`
	Shipment       bool   `json:"shipment"`
	Location_id    uint32 `json:"location_id"`
}

type Account struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       role   `json:"role"`
	Address_id uint32 `json:"address_id"`
}
type Accounts []*Account
