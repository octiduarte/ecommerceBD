package entities

type ProductImage struct {
	ProductImageID int64  `json:"product_image_id"`
	URL            string `json:"url"`
	ProductID      int64  `json:"product_id"`
}
