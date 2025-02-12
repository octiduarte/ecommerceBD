package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"simi/internal/rest/handlers"
)

func InitRoutes(productsHandler handlers.ProductsHandler, storesHandler handlers.StoresHandler,
	mainStoresHandler handlers.MainStoresHandler, productsImageHandler handlers.ImageHandler) *mux.Router {

	r := mux.NewRouter()

	// MAIN
	r.HandleFunc("/main/stores/{store_id:[0-9]+}", mainStoresHandler.GetMainStores).
		Methods(http.MethodGet)

	// Image
	r.HandleFunc("/image/product", productsImageHandler.UploadProductImage).
		Methods(http.MethodPost)
	r.HandleFunc("/image/store", productsImageHandler.UploadStoreImage).
		Methods(http.MethodPost)

	// CRUD Stores
	r.HandleFunc("/stores", storesHandler.GetStores).
		Methods(http.MethodGet)
	r.HandleFunc("/stores/{id:[0-9]+}", storesHandler.GetStoreByID).
		Methods(http.MethodGet)
	r.HandleFunc("/stores", storesHandler.SetStores).
		Methods(http.MethodPost)

	// CRUD Products
	r.HandleFunc("/products/{id:[0-9]+}", productsHandler.GetProductsByID).
		Methods(http.MethodGet)
	r.HandleFunc("/products", productsHandler.SetProducts).
		Methods(http.MethodPost)

	return r
}
