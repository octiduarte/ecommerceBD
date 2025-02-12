package interfaces

import (
	"mime/multipart"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

type ProductsService interface {
	GetProductByID(id int) (productResponse model.ProductPage, err error)
	SetProducts(newProduct entities.Product) (productsResponse []entities.Product, err error)
}

type StoresService interface {
	GetStores() (storesResponse []entities.Store, err error)
	GetStoreByID(id int) (storeResponse entities.Store, err error)
	SetStores(newStore entities.Store) (storeID int64, err error)
}

type MainStoresService interface {
	GetMainStoreByID(storeID int) (mainStoreResponse model.MainResponse, err error)
}

type ImageService interface {
	UploadProductsImage(file multipart.File, handler *multipart.FileHeader, productID string) error
	UploadStoreImage(fileLogo multipart.File, handlerLogo *multipart.FileHeader, fileBanner multipart.File,
		handlerBanner *multipart.FileHeader, StoreID string) error
}
