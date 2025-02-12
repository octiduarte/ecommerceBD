package entities

type Discount struct {
	DiscountID     int    `json:"discount_id"`
	DiscountType   string `json:"discount_type"` //percent|fixed|installments
	DiscountAmount int    `json:"discount_amount"`
}
