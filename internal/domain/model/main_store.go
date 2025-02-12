package model

import "simi/internal/domain/model/entities"

type MainResponse struct {
	Store      MainStore     `json:"store"`
	Categories []string      `json:"categories"`
	Products   []MainProduct `json:"products"`
}

type MainStore struct {
	Name        string                 `json:"name"`
	Logo        string                 `json:"logo"`
	Banner      string                 `json:"banner"`
	SocialMedia []entities.SocialMedia `json:"socialMedia"`
}

type MainProduct struct {
	ProductID       int64   `json:"product_id"`
	Image           string  `json:"image"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
	MainDescription string  `json:"main_description"`
}
