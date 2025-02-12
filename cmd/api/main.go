package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"simi/cmd/api/dependencies"
	"simi/cmd/api/routes"
	"simi/internal/utils/db"
)

func main() {
	database, err := db.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error: ", err)
		panic(err)
	}
	defer database.Close()

	fmt.Println("Connected to database")

	dm := dependencies.NewDependencyManager(database)

	productsHandlers := dm.ProductsHandler()
	storesHandler := dm.StoresHandler()
	mainStoresHandler := dm.MainStoresHandler()
	productsImageHandler := dm.ProductsImageHandler()

	r := routes.InitRoutes(productsHandlers, storesHandler, mainStoresHandler, productsImageHandler)

	// Configurar CORS para permitir solicitudes desde localhost:3000
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	fs := http.FileServer(http.Dir("./internal/image"))
	r.PathPrefix("/image/").Handler(http.StripPrefix("/image/", fs))

	err = http.ListenAndServe(":8000", corsHandler)

	if err != nil {
		fmt.Println("Error starting server with error: ", err)
		panic(err)
	}

}
