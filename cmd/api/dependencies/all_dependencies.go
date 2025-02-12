package dependencies

import (
	"database/sql"
	"simi/internal/domain/interfaces"
	"simi/internal/domain/services/image_service"
	"simi/internal/domain/services/main_stores_service"
	"simi/internal/domain/services/products_service"
	"simi/internal/domain/services/stores_service"
	"simi/internal/repositories/category_repo"
	"simi/internal/repositories/email_repo"
	"simi/internal/repositories/image_repo"
	"simi/internal/repositories/main_stores_repo"
	"simi/internal/repositories/products_repo"
	"simi/internal/repositories/stores_repo"
	"simi/internal/rest/handlers"
)

// AllDependencies production scopes dependencies
type AllDependencies struct {
	db *sql.DB
}

// Handlers

func (d AllDependencies) ProductsHandler() handlers.ProductsHandler {
	return handlers.NewProductsHandler(d.ProductsService())
}

func (d AllDependencies) StoresHandler() handlers.StoresHandler {
	return handlers.NewStoresHandler(d.StoresService())
}

func (d AllDependencies) MainStoresHandler() handlers.MainStoresHandler {
	return handlers.NewMainStoresHandler(d.MainStoresService())
}

func (d AllDependencies) ProductsImageHandler() handlers.ImageHandler {
	return handlers.NewProductsImageHandler(d.ImageService())
}

// Services

func (d AllDependencies) ProductsService() interfaces.ProductsService {
	return products_service.NewProductsService(d.ProductsRepository())
}

func (d AllDependencies) StoresService() interfaces.StoresService {
	return stores_service.NewStoresService(d.StoresRepository())
}

func (d AllDependencies) MainStoresService() interfaces.MainStoresService {
	return main_stores_service.NewMainStoresService(d.MainStoresRepository())
}

func (d AllDependencies) ImageService() interfaces.ImageService {
	return image_service.NewImageService(d.ImageRepository(), d.ProductsRepository(), d.StoresRepository())
}

// Repositories

func (d AllDependencies) ProductsRepository() interfaces.ProductsRepository {
	return products_repo.NewProductsRepository(d.db)
}

func (d AllDependencies) StoresRepository() interfaces.StoresRepository {
	return stores_repo.NewStoresRepository(d.db)
}

func (d AllDependencies) CategoryRepository() interfaces.CategoriesRepository {
	return category_repo.NewCategoryRepo(d.db)
}

func (d AllDependencies) ImageRepository() interfaces.ImageRepository {
	return image_repo.NewProductsImageRepo(d.db)
}

func (d AllDependencies) EmailRepository() interfaces.EmailRepository {
	return email_repo.NewEmailRepository()
}

func (d AllDependencies) MainStoresRepository() interfaces.MainStoresRepository {
	return main_stores_repo.NewMainStoresRepo(d.db)
}
