package entities

type PaymentMethod struct {
	PaymentMethodID int64  `json:"payment_method_id"`
	StoreID         int64  `json:"store_id"`
	Type            string `json:"type"`
	CVU             string `json:"cvu"`
	Alias           string `json:"alias"`
}
