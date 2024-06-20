package types

type place string
type role string

const (
	work   place = "work"
	home   place = "house"
	user   role  = "user"
	admin  role  = "admin"
	worker role  = "worker"
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
}

type Account struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     role   `json:"role"`
}
type Accounts []*Account
