package model

type Account struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Rols   `json:"role"`
	AddressID uint32 `json:"addressID"`
	CardID    uint32 `json:"cardID"`
}

type Accounts []*Account

type Rols string

const (
	user   Rols = "user"
	admin  Rols = "admin"
	worker Rols = "worker"
)
