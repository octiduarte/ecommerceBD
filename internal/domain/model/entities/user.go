package entities

type User struct {
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Address   string `json:"address"`
}
