package interfaces

import (
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

type ProductsRepository interface {
	SetProducts(product entities.Product) error
	GetProducts() (productsResponse []entities.Product, err error)
	GetProductById(id int) (product model.ProductPage, err error)
	GetProductsWithCategoriesByStoreID(storeID int, limit string) (productsResponse []model.MainProduct, err error)
	GetStoreNameByProductID(productID string) (storeName string, err error)
}

type StoresRepository interface {
	SetStores(store entities.Store) (storeID int64, err error)
	GetStores() (storesResponse []entities.Store, err error)
	GetStoreById(id int) (store entities.Store, err error)
	GetMainStoreById(id int) (store model.MainStore, err error)
	SetStoreImage(storeID int, imageLogo, imageBanner string) error
}

type CategoriesRepository interface {
	GetCategoriesByIDs(categoryIDs []int64) (categories []string, err error)
}

type ImageRepository interface {
	UploadImageToProduct(productID, imagePath string) error
	GetProductsImage(productID int64) (path string, err error)
}

type EmailRepository interface {
	SendEmail(user entities.User, order model.Order) error
}

type MainStoresRepository interface {
	GetMainStore(storeID int) (response model.MainResponse, err error)
}
