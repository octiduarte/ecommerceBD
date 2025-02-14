package products_service

import (
	"simi/internal/domain/interfaces"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

const pathImageAccess = "https://ecommercebd-production-6168.up.railway.app/image/"

type ProductsService struct {
	productsRepository interfaces.ProductsRepository
}

func NewProductsService(productsRepository interfaces.ProductsRepository) ProductsService {
	return ProductsService{productsRepository: productsRepository}
}

func (s ProductsService) GetProductByID(id int) (productResponse model.ProductPage, err error) {
	var images []string
	productResponse, err = s.productsRepository.GetProductById(id)
	if err != nil {
		return productResponse, err
	}
	for _, image := range productResponse.Images {
		images = append(images, pathImageAccess+productResponse.StoreName+image)
	}
	productResponse.Images = images
	return productResponse, nil
}

func (s ProductsService) SetProducts(newProduct entities.Product) (productsResponse []entities.Product, err error) {
	err = s.productsRepository.SetProducts(newProduct)
	if err != nil {
		return productsResponse, err
	}
	return productsResponse, nil
}
