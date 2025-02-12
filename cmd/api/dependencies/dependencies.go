package dependencies

import (
	"database/sql"
	"simi/internal/domain/interfaces"
	"simi/internal/rest/handlers"
)

type DependencyManager interface {

	// Handlers
	ProductsHandler() handlers.ProductsHandler
	StoresHandler() handlers.StoresHandler
	MainStoresHandler() handlers.MainStoresHandler
	ProductsImageHandler() handlers.ImageHandler

	// Services
	ProductsService() interfaces.ProductsService
	StoresService() interfaces.StoresService
	MainStoresService() interfaces.MainStoresService
	ImageService() interfaces.ImageService

	// Repositories
	ProductsRepository() interfaces.ProductsRepository
	StoresRepository() interfaces.StoresRepository
	CategoryRepository() interfaces.CategoriesRepository
	ImageRepository() interfaces.ImageRepository
	EmailRepository() interfaces.EmailRepository
	MainStoresRepository() interfaces.MainStoresRepository
}

// NewDependencyManager create a new dependency manager based in the scope
func NewDependencyManager(db *sql.DB) DependencyManager {
	return AllDependencies{db: db}
}
