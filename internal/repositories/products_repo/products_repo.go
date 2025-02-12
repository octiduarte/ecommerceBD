package products_repo

import (
	"database/sql"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

type ProductsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
	return ProductsRepository{db: db}
}

func (r ProductsRepository) SetProducts(product entities.Product) error {
	_, err := r.db.Exec("INSERT INTO Product (product_id, name, main_description, price, category_id, store_id, discount_id) VALUES (?, ?, ?)",
		product.ProductID, product.Name, product.MainDescription, product.Price, product.CategoryID, product.StoreID, product.DiscountID)
	if err != nil {
		return err
	}
	return nil
}

func (r ProductsRepository) GetProducts() (productsResponse []entities.Product, err error) {
	rows, err := r.db.Query("SELECT product_id, name, main_description, price, category_id, store_id, discount_id FROM Product")
	if err != nil {
		return productsResponse, err
	}

	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.ProductID, &product.Name, &product.MainDescription, &product.Price, &product.CategoryID, &product.StoreID, &product.DiscountID)
		if err != nil {
			return productsResponse, err
		}
		productsResponse = append(productsResponse, product)
	}

	return productsResponse, nil
}

func (r ProductsRepository) GetProductById(id int) (product model.ProductPage, err error) {
	var productPage model.ProductPage
	productPage.Images = []string{}
	productPage.Sizes = []string{}
	productPage.Colors = []string{}

	// Usar mapas para evitar duplicados
	sizeSet := make(map[string]struct{})
	colorSet := make(map[string]struct{})
	imageSet := make(map[string]struct{})

	rows, err := r.db.Query("SELECT st.Name, p.Product_id, p.Name, p.Main_description, p.Long_description, p.Price, "+
		"c.Category_name, COALESCE(d.Type, 'No Discount') AS Discount_Type, COALESCE(d.Amount, 0) AS Discount_Amount, pd.Stock_count, s.Size AS Size, "+
		"co.Color AS Color, pi.Url AS Product_Image FROM Product p LEFT JOIN Category c ON p.Category_id = c.Category_id LEFT JOIN Discount d ON p.Discount_id = d.Discount_id "+
		"LEFT JOIN Product_detail pd ON p.Product_id = pd.Product_id LEFT JOIN Size s ON pd.Size_id = s.Size_id LEFT JOIN Color co ON pd.Color_id = co.Color_id LEFT JOIN "+
		"Product_image pi ON p.Product_id = pi.Product_id LEFT JOIN Store st ON p.store_id = st.store_id WHERE p.Product_id = ?;", id)
	if err != nil {
		return productPage, err
	}
	defer rows.Close()

	for rows.Next() {
		var size, color, image, longDescription sql.NullString
		err := rows.Scan(
			&productPage.StoreName,
			&productPage.ProductID,
			&productPage.Name,
			&productPage.MainDescription,
			&longDescription,
			&productPage.Price,
			&productPage.CategoryName,
			&productPage.DiscountType,
			&productPage.DiscountAmount,
			&productPage.StockCount,
			&size,
			&color,
			&image,
		)
		if err != nil {
			return productPage, err
		}

		if longDescription.Valid {
			productPage.LongDescription = longDescription.String
		}

		// Añadir tallas sin duplicados
		if size.Valid {
			if _, exists := sizeSet[size.String]; !exists {
				productPage.Sizes = append(productPage.Sizes, size.String)
				sizeSet[size.String] = struct{}{}
			}
		}

		// Añadir colores sin duplicados
		if color.Valid {
			if _, exists := colorSet[color.String]; !exists {
				productPage.Colors = append(productPage.Colors, color.String)
				colorSet[color.String] = struct{}{}
			}
		}

		// Añadir imágenes sin duplicados
		if image.Valid {
			if _, exists := imageSet[image.String]; !exists {
				productPage.Images = append(productPage.Images, image.String)
				imageSet[image.String] = struct{}{}
			}
		}
	}

	return productPage, nil
}

func (r ProductsRepository) GetProductsWithCategoriesByStoreID(storeID int, limit string) (productsResponse []model.MainProduct, err error) {
	var rows *sql.Rows

	if limit != "" {
		rows, err = r.db.Query("SELECT p.product_id, p.name, p.price, c.category_name, p.main_description FROM Product as p JOIN Category as c ON (p.category_id = c.category_id) where store_id = ? limit ?", storeID, limit)
		if err != nil {
			return productsResponse, err
		}
	} else {
		rows, err = r.db.Query("SELECT p.product_id, p.name, p.price, c.category_name, p.main_description FROM Product as p JOIN Category as c ON (p.category_id = c.category_id) where store_id = ?", storeID)
		if err != nil {
			return productsResponse, err
		}
	}

	for rows.Next() {
		var product model.MainProduct
		err = rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Category, &product.MainDescription)
		if err != nil {
			return productsResponse, err
		}
		productsResponse = append(productsResponse, product)
	}

	return productsResponse, nil
}

func (r ProductsRepository) GetStoreNameByProductID(productID string) (storeName string, err error) {
	row := r.db.QueryRow("SELECT s.name FROM Store as s JOIN Product as p ON (s.store_id = p.store_id) where p.product_id = ? limit 1", productID)
	err = row.Scan(&storeName)
	if err != nil {
		return storeName, err
	}
	return storeName, nil
}
