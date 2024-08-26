package types

type role string

const (
	// Role
	user   role = "user"
	admin  role = "admin"
	worker role = "worker"
)

type Account struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     role   `json:"role"`
}
type Accounts []*Account
