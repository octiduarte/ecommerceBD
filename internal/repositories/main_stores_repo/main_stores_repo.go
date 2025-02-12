package main_stores_repo

import (
	"database/sql"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

const pathImageAccess = "http://localhost:8000/image/"

type MainStoresRepo struct {
	db *sql.DB
}

func NewMainStoresRepo(db *sql.DB) MainStoresRepo {
	return MainStoresRepo{db: db}
}

func (r MainStoresRepo) GetMainStore(storeID int) (response model.MainResponse, err error) {
	query := `
	SELECT s.Name as store_name, s.Logo as store_logo, s.Banner as store_banner,
    		sm.Name as social_media_name, sm.Url as social_media_url,
    		c.Category_name, p.Product_id, p.Name as product_name, p.Price,
    		p.Main_description, pi.Url as product_image
	FROM Store s
	LEFT JOIN Social_media sm ON s.Store_id = sm.Store_id
	JOIN Product p ON s.Store_id = p.Store_id
	JOIN Product_image pi ON p.Product_id = pi.Product_id
	JOIN Category c ON p.Category_id = c.Category_id
	WHERE s.Store_id = ?;
	`

	rows, err := r.db.Query(query, storeID)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	socialMediaMap := make(map[string]string)
	categoryMap := make(map[string]bool)
	productMap := make(map[int]bool)

	for rows.Next() {
		var (
			socialName, socialURL, categoryName, productName, productDesc, productImage string
			productID, price                                                            int
		)
		err := rows.Scan(&response.Store.Name, &response.Store.Logo, &response.Store.Banner, &socialName, &socialURL,
			&categoryName, &productID, &productName, &price, &productDesc, &productImage)
		if err != nil {
			return response, err
		}

		if socialName != "" {
			if _, exists := socialMediaMap[socialName]; !exists {
				socialMedia := entities.SocialMedia{Name: socialName, URL: socialURL}
				response.Store.SocialMedia = append(response.Store.SocialMedia, socialMedia)
				socialMediaMap[socialName] = socialURL
			}
		}

		if !categoryMap[categoryName] {
			response.Categories = append(response.Categories, categoryName)
			categoryMap[categoryName] = true
		}

		if !productMap[productID] {
			product := model.MainProduct{ProductID: int64(productID),
				Image:           pathImageAccess + response.Store.Name + "/" + productImage,
				Name:            productName,
				Price:           float64(price),
				Category:        categoryName,
				MainDescription: productDesc}
			response.Products = append(response.Products, product)
			productMap[productID] = true
		}
	}

	response.Store.Logo = pathImageAccess + response.Store.Name + "/" + response.Store.Logo
	response.Store.Banner = pathImageAccess + response.Store.Name + "/" + response.Store.Banner

	return response, nil
}
