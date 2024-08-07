package types

type place string
type role string

const (
	// Role
	user   role = "user"
	admin  role = "admin"
	worker role = "worker"

	// Place
	work place = "work"
	home place = "house"
)

type Address struct {
	ID             uint32 `json:"id"`
	Place          place  `json:"place"`
	Street         string `json:"street"`
	Height         uint16 `json:"height"`
	Floor          uint16 `json:"floor"`
	Department     uint16 `json:"department"`
	BetweenStreets string `json:"betweenStreets"`
	Observations   string `json:"observations"`
	Shipment       bool   `json:"shipment"`
	Location_ID    uint32 `json:"location_id"`
	Account_ID     uint32 `json:"account_id"`
}
type Addresses []*Address

type Account struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     role   `json:"role"`
}
type Accounts []*Account
