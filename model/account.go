package model

type Account struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      rols   `json:"role"`
	AddressID uint32 `json:"addressID"`
	CardID    uint32 `json:"cardID"`
}

type Accounts []*Account

type rols string

const (
	user   rols = "user"
	admin  rols = "admin"
	worker rols = "worker"
)
