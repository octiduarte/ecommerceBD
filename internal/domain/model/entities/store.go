package entities

type Store struct {
	StoreID   int64  `json:"store_id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Banner    string `json:"banner"`
	Address   string `json:"address"`
	Cellphone string `json:"cellphone"`
}
