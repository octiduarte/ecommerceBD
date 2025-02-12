package image_repo

import (
	"database/sql"
	"simi/internal/domain/interfaces"
)

type ProductsImageRepo struct {
	db              *sql.DB
	productsService interfaces.ProductsService
}

func NewProductsImageRepo(db *sql.DB) ProductsImageRepo {
	return ProductsImageRepo{db: db}
}

func (r ProductsImageRepo) UploadImageToProduct(productID, imagePath string) error {
	_, err := r.db.Exec("INSERT INTO Product_image (product_id, url) VALUES (?, ?)", productID, imagePath)
	if err != nil {
		return err
	}
	return err
}

func (r ProductsImageRepo) GetProductsImage(productID int64) (path string, err error) {
	row := r.db.QueryRow("SELECT url FROM Product_image WHERE product_id = ?", productID)
	err = row.Scan(&path)
	if err != nil {
		return path, err
	}
	return path, nil
}
