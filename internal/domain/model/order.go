package model

import "simi/internal/domain/model/entities"

type Order struct {
	Products   []entities.Product `json:"products"`
	TotalPrice float64
}
